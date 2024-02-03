package data

import (
	cd "MyEcommerce/features/cart/data"
	"MyEcommerce/features/order"

	ud "MyEcommerce/features/user/data"

	"gorm.io/gorm"
)

type Order struct {
	ID string `gorm:"type:varchar(36);primary_key" json:"id"`
	gorm.Model
	UserID      uint
	User        ud.User
	Address     string
	PaymentType string
	GrossAmount int
	Status      string
	Bank        string
	VaNumber    int
}

type OrderItem struct {
	gorm.Model
	OrderID string
	Order   Order
	CartID  uint
	Cart    cd.Cart
}

func CoreToModel(reqNotif order.OrderCore) Order {
	return Order{
		Status: reqNotif.Status,
	}
}

func CoreToModelOrder(input order.OrderCore) Order {
	return Order{
		ID:          input.ID,
		UserID:      input.UserID,
		Address:     input.Address,
		GrossAmount: input.GrossAmount,
		Status:      input.Status,
		VaNumber:    input.VaNumber,
		Bank:        input.Bank,
	}
}

func CoreToModelOrderCancle(input order.OrderCore) Order {
	return Order{
		Status: input.Status,
	}
}

func (o Order) ModelToCoreOrderUser() order.OrderCore {
	return order.OrderCore{
		Status:      o.Status,
		GrossAmount: o.GrossAmount,
		VaNumber:    o.VaNumber,
		Bank:        o.Bank,
	}
}

func (o Order) ModelToCoreOrderAdmin() order.OrderCore {
	return order.OrderCore{
		// ID:          o.ID,
		CreatedAt:   o.CreatedAt,
		Bank:        o.Bank,
		GrossAmount: o.GrossAmount,
		Address:     o.Address,
		Status:      o.Status,
	}
}

func (ot OrderItem) ModelToCoreOrderItemUser() order.OrderItemCore {
	return order.OrderItemCore{
		OrderID: ot.OrderID,
		CartID:  ot.CartID,
		Cart:    ot.Cart.ModelToCore(),
		Order:   ot.Order.ModelToCoreOrderUser(),
	}
}

func (ot OrderItem) ModelToCoreOrderItemAdmin() order.OrderItemCore {
	return order.OrderItemCore{
		OrderID: ot.OrderID,
		CartID:  ot.CartID,
		Cart:    ot.Cart.ModelToCore(),
		Order:   ot.Order.ModelToCoreOrderAdmin(),
	}
}
