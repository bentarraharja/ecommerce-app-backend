package data

import (
	cd "MyEcommerce/features/cart/data"
	"MyEcommerce/features/order"
	"MyEcommerce/features/user/data"
	"MyEcommerce/utils/externalapi"
	"errors"
	"log"

	"gorm.io/gorm"
)

type orderQuery struct {
	db              *gorm.DB
	paymentMidtrans externalapi.MidtransInterface
}

func New(db *gorm.DB, mi externalapi.MidtransInterface) order.OrderDataInterface {
	return &orderQuery{
		db:              db,
		paymentMidtrans: mi,
	}
}

// InsertOrder implements order.OrderDataInterface.
func (repo *orderQuery) InsertOrder(userIdLogin int, cartIds []uint, inputOrder order.OrderCore) (*order.OrderCore, error) {

	var totalHargaKeseluruhan int
	for _, cartId := range cartIds {
		var cartGorm cd.Cart
		ts := repo.db.Preload("Product").Where("user_id = ? AND id = ?", userIdLogin, cartId).First(&cartGorm)
		if ts.Error != nil {
			return nil, ts.Error
		}
		subTotal := cartGorm.Product.Price * cartGorm.Quantity
		totalHargaKeseluruhan += subTotal
	}

	inputOrder.GrossAmount = totalHargaKeseluruhan

	payment, errPay := repo.paymentMidtrans.NewOrderPayment(inputOrder)
	if errPay != nil {
		return nil, errPay
	}

	orderModel := CoreToModelOrder(inputOrder)
	orderModel.UserID = uint(userIdLogin)

	orderModel.PaymentType = payment.PaymentType
	orderModel.Status = payment.Status
	orderModel.VaNumber = payment.VaNumber

	tx := repo.db.Create(&orderModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	inputOrder.ID = orderModel.ID

	for _, cartId := range cartIds {
		orderItem := OrderItem{
			OrderID: orderModel.ID,
			CartID:  cartId,
		}

		if err := repo.db.Create(&orderItem).Error; err != nil {
			return nil, err
		}

		var cart cd.Cart

		if err := repo.db.Preload("Product").Where("id = ?", cartId).First(&cart).Error; err != nil {
			return nil, err
		}

		cart.Product.Stock -= cart.Quantity

		if err := repo.db.Save(&cart.Product).Error; err != nil {
			return nil, err
		}
	}

	for _, cartId := range cartIds {
		var cartGorm cd.Cart
		ts := repo.db.Where("user_id = ? AND id = ?", userIdLogin, cartId).First(&cartGorm)
		if ts.Error != nil {
			return nil, ts.Error
		}

		td := repo.db.Delete(&cartGorm)
		if td.Error != nil {
			return nil, td.Error
		}

		if td.RowsAffected == 0 {
			return nil, errors.New("error record not found")
		}
	}

	return payment, nil
}

// SelectOrderUser implements order.OrderDataInterface.
func (repo *orderQuery) SelectOrderUser(userIdLogin int) ([]order.OrderItemCore, error) {
	var orderItems []OrderItem
	err := repo.db.Unscoped().Joins("Order").Preload("Cart").Preload("Cart.Product").Preload("Cart.Product.User").Where("user_id = ?", userIdLogin).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}

	orderItemCores := make([]order.OrderItemCore, len(orderItems))
	for i, item := range orderItems {
		orderItemCores[i] = item.ModelToCoreOrderItemUser()
	}

	return orderItemCores, nil
}

// SelectOrderAdmin implements order.OrderDataInterface.
func (repo *orderQuery) SelectOrderAdmin(userIdLogin, page, limit int) ([]order.OrderItemCore, int, error) {
	var userDataGorm data.User
	tx := repo.db.Where("role = 'admin' AND id = ?", userIdLogin).First(&userDataGorm, userIdLogin)
	if tx.Error != nil {
		return nil, 0, errors.New("Sorry, your role does not have this access.")
	}

	var orderItems []OrderItem
	var totalData int64
	tc := repo.db.Unscoped().Model(&orderItems).Count(&totalData)
	if tc.Error != nil {
		return nil, 0, tc.Error
	}
	totalPage := int((totalData + int64(limit) - 1) / int64(limit))
	log.Println(totalPage)

	err := repo.db.Unscoped().Joins("Order").Preload("Cart").Preload("Cart.Product").Limit(limit).Offset((page - 1) * limit).Find(&orderItems).Error
	if err != nil {
		return nil, 0, err
	}

	orderItemCores := make([]order.OrderItemCore, len(orderItems))
	for i, item := range orderItems {
		orderItemCores[i] = item.ModelToCoreOrderItemAdmin()
	}

	return orderItemCores, int(totalPage), nil
}

// SelectOrderAdmin implements order.OrderDataInterface.
func (repo *orderQuery) CancleOrder(userIdLogin int, orderId string, orderCore order.OrderCore) error {
	if orderCore.Status == "cancelled" {
		repo.paymentMidtrans.CancelOrderPayment(orderId)
	}

	dataGorm := CoreToModelOrderCancle(orderCore)
	tx := repo.db.Model(&Order{}).Where("id = ? AND user_id = ?", orderId, userIdLogin).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Update implements user.UserDataInterface.
func (repo *orderQuery) WebhoocksData(reqNotif order.OrderCore) error {
	dataGorm := CoreToModel(reqNotif)
	tx := repo.db.Model(&Order{}).Where("id = ?", reqNotif.ID).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}
