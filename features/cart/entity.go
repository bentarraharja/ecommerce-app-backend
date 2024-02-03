package cart

import (
	"MyEcommerce/features/product"
	"MyEcommerce/features/user"
	"time"
)

type Core struct {
	ID        uint
	Quantity  int
	UserID    uint
	ProductID uint
	User      user.Core
	Product   product.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartDataInterface interface {
	Insert(userIdLogin, productId int) error
	Select(userIdLogin int) ([]Core, error)
	Update(userIdLogin, cartId int, input Core) error
	Delete(userIdLogin, cartId int) error
}

// interface untuk Service Layer
type CartServiceInterface interface {
	Create(userIdLogin, productId int) error
	Get(userIdLogin int) ([]Core, error)
	UpdateCart(userIdLogin int, cartId int, input Core) error
	DeleteCart(userIdLogin int, cartId int) error
}
