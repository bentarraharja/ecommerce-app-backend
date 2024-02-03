package service

import (
	"MyEcommerce/features/order"
	"MyEcommerce/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := New(repo)

	cartIds := []uint{1, 2, 3}
	inputOrder := order.OrderCore{
		Address: "JAKARTA BARAT UTARA DAYA SELATAN",
		Bank:    "BCA",
	}

	t.Run("error when cartIds is empty", func(t *testing.T) {
		idBarang := cartIds
		cartIds = []uint{}
		result, err := srv.CreateOrder(1, cartIds, inputOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "masukan barang anda")

		cartIds = idBarang
	})

	t.Run("error when address is empty", func(t *testing.T) {
		alamatAsli := inputOrder.Address
		inputOrder.Address = ""
		result, err := srv.CreateOrder(1, cartIds, inputOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "masukan alamat anda")

		inputOrder.Address = alamatAsli
	})

	t.Run("error from repository on create", func(t *testing.T) {
		repo.On("InsertOrder", 1, cartIds, inputOrder).Return(nil, errors.New("database error")).Once()

		result, err := srv.CreateOrder(1, cartIds, inputOrder)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		returnData := &order.OrderCore{
			ID: "asdhabd",
		}
		repo.On("InsertOrder", 1, cartIds, inputOrder).Return(returnData, nil).Once()

		result, err := srv.CreateOrder(1, cartIds, inputOrder)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestGetOrderUser(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := New(repo)

	returnData := []order.OrderItemCore{
		{
			CartID:  1,
			OrderID: "abdksjc",
		},
	}

	t.Run("error from repository on get", func(t *testing.T) {
		repo.On("SelectOrderUser", 1).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetOrderUser(1)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("SelectOrderUser", 1).Return(returnData, nil).Once()

		result, err := srv.GetOrderUser(1)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repo.AssertExpectations(t)
	})
}

func TestGetOrderAdmin(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := New(repo)

	returnData := []order.OrderItemCore{
		{
			OrderID: "1",
			CartID:  1,
		},
	}

	t.Run("default page and limit", func(t *testing.T) {
		userIdLogin := 1
		page := 0
		limit := 0

		repo.On("SelectOrderAdmin", 1, 1, 10).Return(returnData, 0, nil).Once()

		result, _, err := srv.GetOrderAdmin(userIdLogin, page, limit)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		page := 1
		limit := 10
		userIdLogin := 1

		repo.On("SelectOrderAdmin", userIdLogin, page, limit).Return(nil, 0, errors.New("database error")).Once()

		result, _, err := srv.GetOrderAdmin(userIdLogin, page, limit)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		page := 1
		limit := 10
		userIdLogin := 1

		repo.On("SelectOrderAdmin", userIdLogin, page, limit).Return(returnData, 0, nil).Once()

		result, _, err := srv.GetOrderAdmin(userIdLogin, page, limit)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})
}

func TestCancleOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := New(repo)

	orderCore := order.OrderCore{
		Status: "cancelled",
	}

	t.Run("success", func(t *testing.T) {
		repo.On("CancleOrder", 1, "1", orderCore).Return(nil).Once()

		orderCore.Status = ""
		err := srv.CancleOrder(1, "1", orderCore)
		orderCore.Status = "cancelled"
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestWebhoocksService(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := New(repo)

	reqNotif := order.OrderCore{
		ID: "oreifin",
	}

	t.Run("invalid order id", func(t *testing.T) {
		idOrder := reqNotif.ID
		reqNotif.ID = ""
		err := srv.WebhoocksService(reqNotif)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid order id")

		reqNotif.ID = idOrder
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("WebhoocksData", reqNotif).Return(errors.New("database error")).Once()

		err := srv.WebhoocksService(reqNotif)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("WebhoocksData", reqNotif).Return(nil).Once()

		err := srv.WebhoocksService(reqNotif)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}
