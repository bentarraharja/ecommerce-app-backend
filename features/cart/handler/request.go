package handler

import (
	"MyEcommerce/features/cart"
)

type CartRequest struct {
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity" form:"quantity"`
}

func RequestToCore(input CartRequest) cart.Core {
	return cart.Core{
		Quantity: input.Quantity,
	}
}
