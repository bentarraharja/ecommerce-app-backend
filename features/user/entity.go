package user

import "time"

type Core struct {
	ID           uint
	Name         string `validate:"required"`
	UserName     string `validate:"required"`
	Email        string `validate:"required,email"`
	Password     string `validate:"required"`
	Role         string
	PhotoProfile string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input Core) error
	SelectById(userIdLogin int) (*Core, error)
	Update(userIdLogin int, input Core) error
	Delete(userIdLogin int) error
	Login(email, password string) (data *Core, err error)
	SelectAdminUsers(page, limit int) ([]Core, error, int)
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input Core) error
	GetById(userIdLogin int) (*Core, error)
	Update(userIdLogin int, input Core) error
	Delete(userIdLogin int) error
	Login(email, password string) (data *Core, token string, err error)
	GetAdminUsers(userIdLogin, page, limit int) ([]Core, error, int)
}
