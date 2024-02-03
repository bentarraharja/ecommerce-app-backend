package handler

import (
	"MyEcommerce/features/order"
	ph "MyEcommerce/features/product/handler"
	uh "MyEcommerce/features/user/handler"
	"time"
)

type OrderResponse struct {
	OrderID string `json:"order_id" form:"order_id"`
	Address string `json:"address" form:"address"`
	Payment PaymentResponse
}

type PaymentResponse struct {
	StatusCode      string `json:"status_code" form:"status_code"`
	StatusMessage   string `json:"status_message" form:"status_message"`
	Status          string `json:"status" form:"status"`
	PaymentType     string `json:"payment_type" form:"payment_type"`
	GrossAmount     int    `json:"gross_amount" form:"gross_amount"`
	Bank            string `json:"bank" form:"bank"`
	VaNumber        int    `json:"va_number" form:"va_number"`
	TransactionId   string `json:"transaction_id" form:"transaction_id"`
	Currency        string `json:"currency" form:"currency"`
	TransactionTime string `json:"transaction_time" form:"transaction_time"`
	FraudStatus     string `json:"fraud_status" form:"fraud_status"`
	ExpiredAt       string `json:"expired_at" form:"expired_at"`
}

type OrderUserItemResponse struct {
	Product     ph.CartProductResponse `json:"product"`
	Quantity    int                    `json:"quantity"`
	Status      string                 `json:"status"`
	GrossAmount int                    `json:"gross_amount"`
	VaNumber    int                    `json:"va_number"`
	Bank        string                 `json:"bank"`
}

type GetOrderUserResponse struct {
	ID    string                  `json:"order_id"`
	Order []OrderUserItemResponse `json:"order"`
}

type OrderAdminItemResponse struct {
	OrderID     string                  `json:"order_id"`
	Product     ph.AdminProductResponse `json:"product"`
	Quantity    int                     `json:"quantity"`
	CreatedAt   time.Time               `json:"created_at"`
	Bank        string                  `json:"bank"`
	GrossAmount int                     `json:"gross_amount"`
	Address     string                  `json:"address"`
	Status      string                  `json:"status"`
}

type GetOrderAdminResponse struct {
	Order []OrderAdminItemResponse `json:"order"`
}

func CoreToResponseOrderUser(data order.OrderCore, items []order.OrderItemCore) []GetOrderUserResponse {
	orderMap := make(map[string][]OrderUserItemResponse)
	for _, item := range items {
		user := uh.CartUserResponse{
			Name: item.Cart.Product.User.Name,
		}

		orderItem := OrderUserItemResponse{
			Product: ph.CartProductResponse{
				Name:         item.Cart.Product.Name,
				Price:        item.Cart.Product.Price,
				PhotoProduct: item.Cart.Product.PhotoProduct,
				Toko:         user,
			},
			Quantity:    item.Cart.Quantity,
			Status:      data.Status,
			GrossAmount: data.GrossAmount,
			VaNumber:    data.VaNumber,
			Bank:        data.Bank,
		}
		orderMap[item.OrderID] = append(orderMap[item.OrderID], orderItem)
	}

	var response []GetOrderUserResponse
	for id, orderItems := range orderMap {
		response = append(response, GetOrderUserResponse{
			ID:    id,
			Order: orderItems,
		})
	}
	return response
}

func CoreToResponseOrderAdmin(items []order.OrderItemCore) GetOrderAdminResponse {
	orderItems := make([]OrderAdminItemResponse, len(items))
	for i, item := range items {
		orderItems[i] = OrderAdminItemResponse{
			OrderID: item.OrderID,
			Product: ph.AdminProductResponse{
				Name:  item.Cart.Product.Name,
				Price: item.Cart.Product.Price,
			},
			Quantity:    item.Cart.Quantity,
			CreatedAt:   item.Order.CreatedAt,
			Bank:        item.Order.Bank,
			GrossAmount: item.Order.GrossAmount,
			Address:     item.Order.Address,
			Status:      item.Order.Status,
		}
	}

	return GetOrderAdminResponse{
		Order: orderItems,
	}
}

func CoreToResponse(data *order.OrderCore) OrderResponse {
	var result = OrderResponse{
		OrderID: data.ID,
		Address: data.Address,
		Payment: PaymentResponse{
			StatusCode:      data.Payment.StatusCode,
			StatusMessage:   data.Payment.StatusMessage,
			Status:          data.Status,
			PaymentType:     data.PaymentType,
			GrossAmount:     data.GrossAmount,
			Bank:            data.Bank,
			VaNumber:        data.VaNumber,
			TransactionId:   data.Payment.TransactionId,
			Currency:        data.Payment.Currency,
			TransactionTime: data.Payment.TransactionTime,
			FraudStatus:     data.Payment.FraudStatus,
			ExpiredAt:       data.Payment.ExpiredAt,
		},
	}
	return result
}
