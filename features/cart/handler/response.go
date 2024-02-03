package handler

import (
	"MyEcommerce/features/cart"
	ph "MyEcommerce/features/product/handler"
	uh "MyEcommerce/features/user/handler"
)

type CartResponse struct {
	ID       uint `json:"id" form:"id"`
	Quantity int  `json:"quantity" form:"quantity"`
	Products ph.CartProductResponse
}

func CoreToResponse(data cart.Core) CartResponse {
	userResponse := uh.CartUserResponse{
		Name: data.Product.User.Name,
	}

	productResponse := ph.CartProductResponse{
		Name:         data.Product.Name,
		Price:        data.Product.Price,
		Stock:        data.Product.Stock,
		PhotoProduct: data.Product.PhotoProduct,
		Toko:         userResponse,
	}

	return CartResponse{
		ID:       data.ID,
		Quantity: data.Quantity,
		Products: productResponse,
	}
}

func CoreToResponseList(data []cart.Core) []CartResponse {
	var results []CartResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
