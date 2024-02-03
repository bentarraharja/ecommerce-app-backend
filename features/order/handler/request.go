package handler

import (
	"MyEcommerce/features/order"

	"github.com/google/uuid"
)

type OrderRequest struct {
	ID          string
	Address     string `json:"address" form:"address"`
	GrossAmount int    `json:"gross_amount" form:"gross_amount"`
	Bank        string `json:"bank" form:"bank"`
	CartIDs     []uint `json:"cart_ids" form:"cart_ids"`
}

type CancleOrderRequest struct {
	Status string `json:"status"`
}

type WebhoocksRequest struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
}

func RequestToCoreOrder(input OrderRequest) order.OrderCore {
	return order.OrderCore{
		ID:          uuid.New().String(),
		Address:     input.Address,
		GrossAmount: input.GrossAmount,
		Bank:        input.Bank,
	}
}

func CancleRequestToCoreOrder(input CancleOrderRequest) order.OrderCore {
	return order.OrderCore{
		Status: input.Status,
	}
}

func WebhoocksRequestToCore(input WebhoocksRequest) order.OrderCore {
	return order.OrderCore{
		ID:     input.OrderID,
		Status: input.TransactionStatus,
	}
}
