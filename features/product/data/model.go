package data

import (
	"MyEcommerce/features/product"
	"MyEcommerce/features/user/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string 
	Description  string 
	Category     string 
	Stock        int    
	Price        int    
	PhotoProduct string 
	UserID       uint
	User         data.User
}

func CoreToModel(input product.Core) Product {
	return Product{
		UserID:       input.UserID,
		Name:         input.Name,
		Description:  input.Description,
		Category:     input.Category,
		Stock:        input.Stock,
		Price:        input.Price,
		PhotoProduct: input.PhotoProduct,
	}
}

func (p Product) ModelToCore() product.Core {
	return product.Core{
		UserID:       p.UserID,
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		Category:     p.Category,
		Stock:        p.Stock,
		Price:        p.Price,
		PhotoProduct: p.PhotoProduct,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		User:         p.User.ModelToCore(),
	}
}
