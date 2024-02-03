package service

import (
	"MyEcommerce/features/order"
	"errors"
)

type orderService struct {
	orderData order.OrderDataInterface
}

func New(repo order.OrderDataInterface) order.OrderServiceInterface {
	return &orderService{
		orderData: repo,
	}
}

func (os *orderService) CreateOrder(userIdLogin int, cartIds []uint, inputOrder order.OrderCore) (*order.OrderCore, error) {
	if len(cartIds) == 0 {
		return nil, errors.New("masukan barang anda")
	}
	if inputOrder.Address == "" {
		return nil, errors.New("masukan alamat anda")
	}

	payment, err := os.orderData.InsertOrder(userIdLogin, cartIds, inputOrder)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// GetOrderUser implements order.OrderServiceInterface.
func (os *orderService) GetOrderUser(userIdLogin int) ([]order.OrderItemCore, error) {
	result, err := os.orderData.SelectOrderUser(userIdLogin)
	return result, err
}

// GetOrderAdmin implements order.OrderServiceInterface.
func (os *orderService) GetOrderAdmin(userIdLogin, page, limit int) ([]order.OrderItemCore, int, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	result, totalPage, err := os.orderData.SelectOrderAdmin(userIdLogin, page, limit)
	return result, totalPage, err
}

// CancleOrder implements order.OrderServiceInterface.
func (os *orderService) CancleOrder(userIdLogin int, orderId string, orderCore order.OrderCore) error {
	if orderCore.Status == "" {
		orderCore.Status = "cancelled"
	}

	err := os.orderData.CancleOrder(userIdLogin, orderId, orderCore)
	return err
}

// WebhoocksService implements order.OrderServiceInterface.
func (os *orderService) WebhoocksService(reqNotif order.OrderCore) error {
	if reqNotif.ID == "" {
		return errors.New("invalid order id")
	}

	err := os.orderData.WebhoocksData(reqNotif)
	if err != nil {
		return err
	}

	return nil
}
