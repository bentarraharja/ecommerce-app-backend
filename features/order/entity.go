package order

import (
	cd "MyEcommerce/features/cart"
	ud "MyEcommerce/features/user"
	"time"
)

type OrderCore struct {
	// ID uint
	ID          string
	UserID      uint
	Address     string
	PaymentType string
	GrossAmount int
	Status      string
	VaNumber    int
	Bank        string
	User        ud.Core
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Payment     Payment
	OrderItems  []OrderItemCore
}

type OrderItemCore struct {
	OrderID   string
	CartID    uint
	Order     OrderCore
	Cart      cd.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment struct {
	StatusCode      string
	StatusMessage   string
	TransactionId   string
	Currency        string
	TransactionTime string
	FraudStatus     string
	ExpiredAt       string
}

// type Webhoocks struct {
// 	TransactionID     string
// 	StatusCode        string
// 	StatusMessage     string
// 	TransactionStatus string
// 	PaymentType       string
// 	GrossAmount       string
// 	Bank              string
// 	VaNumber          string
// 	Currency          string
// 	TransactionTime   string
// 	FraudStatus       string
// }

// interface untuk Data Layer
type OrderDataInterface interface {
	InsertOrder(userIdLogin int, cartIds []uint, inputOrder OrderCore) (*OrderCore, error)
	SelectOrderUser(userIdLogin int) ([]OrderItemCore, error)
	SelectOrderAdmin(userIdLogin, page, limit int) ([]OrderItemCore, int, error)
	CancleOrder(userIdLogin int, orderId string, orderCore OrderCore) error
	WebhoocksData(reqNotif OrderCore) error
}

// interface untuk Service Layer
type OrderServiceInterface interface {
	CreateOrder(userIdLogin int, cartIds []uint, inputOrder OrderCore) (*OrderCore, error)
	GetOrderUser(userIdLogin int) ([]OrderItemCore, error)
	GetOrderAdmin(userIdLogin, page, limit int) ([]OrderItemCore, int, error)
	CancleOrder(userIdLogin int, orderId string, orderCore OrderCore) error
	WebhoocksService(reqNotif OrderCore) error
}
