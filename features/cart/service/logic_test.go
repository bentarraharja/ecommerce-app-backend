package service

import (
	"MyEcommerce/features/cart"
	"MyEcommerce/features/product"
	"MyEcommerce/features/user"
	"MyEcommerce/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.CartData)
	srv := New(repo)

	t.Run("error from repository", func(t *testing.T) {
		repo.On("Insert", 1, 1).Return(errors.New("database error")).Once()

		err := srv.Create(1, 1)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Insert", 1, 1).Return(nil).Once()

		err := srv.Create(1, 1)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeleteCart(t *testing.T) {
	repo := new(mocks.CartData)
	srv := New(repo)
	t.Run("error from repository on delete", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("database error")).Once()

		err := srv.DeleteCart(1, 1)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(nil).Once()

		err := srv.DeleteCart(1, 1)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	repo := new(mocks.CartData)
	srv := New(repo)

	returnData := []cart.Core{
		{
			ID:       1,
			Quantity: 1,
			Product: product.Core{
				ID:           1,
				Name:         "Vivo",
				PhotoProduct: "https://res.cloudinary.com/zzsds.jpg",
				User: user.Core{
					ID:   1,
					Name: "lendra",
				},
			},
		},
		{
			ID:       2,
			Quantity: 1,
			Product: product.Core{
				ID:           2,
				Name:         "Samsung",
				PhotoProduct: "https://res.cloudinary.com/zzsds.jpg",
				User: user.Core{
					ID:   1,
					Name: "lendra",
				},
			},
		},
	}

	t.Run("error from repository on select", func(t *testing.T) {
		repo.On("Select", 1).Return(nil, errors.New("database error")).Once()

		result, err := srv.Get(1)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Select", 1).Return(returnData, nil).Once()

		result, err := srv.Get(1)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestUpdateCart(t *testing.T) {
	repo := new(mocks.CartData)
	srv := New(repo)

	returnData := cart.Core{
		Quantity:  1,
	}

	t.Run("error from repository on update", func(t *testing.T) {
		repo.On("Update", 1, 1, returnData).Return(errors.New("database error")).Once()

		err := srv.UpdateCart(1, 1, returnData)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Update", 1, 1, returnData).Return(nil).Once()

		err := srv.UpdateCart(1, 1, returnData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}
