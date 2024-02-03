package service

import (
	"MyEcommerce/features/cart"
)

type cartService struct {
	cartData cart.CartDataInterface
}

// dependency injection
func New(repo cart.CartDataInterface) cart.CartServiceInterface {
	return &cartService{
		cartData: repo,
	}
}

// Create implements cart.CartServiceInterface.
func (cs *cartService) Create(userIdLogin int, productId int) error {
	err := cs.cartData.Insert(userIdLogin, productId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCart implements cart.CartServiceInterface.
func (cs *cartService) DeleteCart(userIdLogin int, cartId int) error {
	err := cs.cartData.Delete(userIdLogin, cartId)
	if err != nil {
		return err
	}
	return nil
}

// Get implements cart.CartServiceInterface.
func (cs *cartService) Get(userIdLogin int) ([]cart.Core, error) {
	result, err := cs.cartData.Select(userIdLogin)
	return result, err
}

// UpdateCart implements cart.CartServiceInterface.
func (cs *cartService) UpdateCart(userIdLogin int, cartId int, input cart.Core) error {
	err := cs.cartData.Update(userIdLogin, cartId, input)
	if err != nil {
		return err
	}
	return nil
}
