package handler

import (
	"MyEcommerce/features/product"
	"MyEcommerce/features/user/handler"
)

type ProductResponse struct {
	ID           uint                        `json:"id" form:"id"`
	Name         string                      `json:"name" form:"name"`
	Description  string                      `json:"description" form:"description"`
	Category     string                      `json:"category" form:"category"`
	Stock        int                         `json:"stock" form:"stock"`
	Price        int                         `json:"price" form:"price"`
	PhotoProduct string                      `json:"photo_product" form:"photo_product"`
	Users        handler.UserProductResponse `json:"toko" form:"toko"`
}

type GetAllProductResponse struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Category     string `json:"category" form:"category"`
	Price        int    `json:"price" form:"price"`
	Stock        int    `json:"stock" form:"stock"`
	PhotoProduct string `json:"photo_product" form:"photo_product"`
}

type CartProductResponse struct {
	Name         string                   `json:"name" form:"name"`
	Price        int                      `json:"price" form:"price"`
	Stock        int                      `json:"stock" form:"stock"`
	PhotoProduct string                   `json:"photo_product" form:"photo_product"`
	Toko         handler.CartUserResponse `json:"toko" form:"toko"`
}

type AdminProductResponse struct {
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}

func CoreToResponse(data product.Core) ProductResponse {
	userResponse := handler.UserProductResponse{
		Name:         data.User.Name,
		UserName:     data.User.UserName,
		PhotoProfile: data.User.PhotoProfile,
	}

	return ProductResponse{
		ID:           data.ID,
		Name:         data.Name,
		Description:  data.Description,
		Category:     data.Category,
		Stock:        data.Stock,
		Price:        data.Price,
		PhotoProduct: data.PhotoProduct,
		Users:        userResponse,
	}
}

func CoreToGetAllResponse(data product.Core) GetAllProductResponse {
	return GetAllProductResponse{
		ID:           data.ID,
		Name:         data.Name,
		Category:     data.Category,
		Price:        data.Price,
		Stock:        data.Stock,
		PhotoProduct: data.PhotoProduct,
	}
}

func CoreToResponseListGetAllProduct(data []product.Core) []GetAllProductResponse {
	var results []GetAllProductResponse
	for _, v := range data {
		results = append(results, CoreToGetAllResponse(v))
	}
	return results
}

func CoreToResponseList(data []product.Core) []ProductResponse {
	var results []ProductResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
