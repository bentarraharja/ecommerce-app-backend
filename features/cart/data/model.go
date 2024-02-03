package data

import (
	"MyEcommerce/features/cart"
	pd "MyEcommerce/features/product/data"
	ud "MyEcommerce/features/user/data"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Quantity  int
	UserID    uint
	ProductID uint
	Product   pd.Product
	User      ud.User
}

func CoreToModel(input cart.Core) Cart {
	return Cart{
		Quantity: input.Quantity,
	}
}

func (c Cart) ModelToCore() cart.Core {
	return cart.Core{
		ID:        c.ID,
		ProductID: c.ProductID,
		Quantity:  c.Quantity,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Product:   c.Product.ModelToCore(),
		User:      c.User.ModelToCore(),
	}
}
