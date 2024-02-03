package service

import (
	"MyEcommerce/features/product"
	"MyEcommerce/features/user"
	"MyEcommerce/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := New(repo)

	returnData := product.Core{
		ID:           1,
		Name:         "samsung",
		Description:  "hp ini murah sekali",
		Category:     "phones",
		Stock:        20,
		Price:        100000,
		PhotoProduct: "https://res.cloudinary.com/ikyft.jpg",
		UserID:       1,
	}

	t.Run("invalid name produk", func(t *testing.T) {
		caseData := returnData
		nameProduk := caseData.Name
		caseData.Name = ""
		err := srv.Create(1, caseData)

		assert.Error(t, err)
		assert.Equal(t, "nama produk tidak boleh kosong", err.Error())

		caseData.Name = nameProduk
	})

	t.Run("invalid price produk", func(t *testing.T) {
		caseData := returnData
		priceProduk := caseData.Price
		caseData.Price = 0
		err := srv.Create(1, caseData)

		assert.Error(t, err)
		assert.Equal(t, "harga produk harus lebih besar dari 0", err.Error())
		caseData.Price = priceProduk
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("Insert", 1, returnData).Return(errors.New("database error")).Once()

		err := srv.Create(1, returnData)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := returnData
		repo.On("Insert", 1, caseData).Return(nil).Once()

		err := srv.Create(1, caseData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := New(repo)

	returnData := []product.Core{
		{
			ID:           1,
			Name:         "samsung",
			Category:     "phones",
			Price:        100000,
			PhotoProduct: "https://res.cloudinary.com/ikyft.jpg",
		},
		{
			ID:           2,
			Name:         "vivo",
			Category:     "phones",
			Price:        120000,
			PhotoProduct: "https://res.cloudinary.com/zzsds.jpg",
		},
	}

	t.Run("default page and limit", func(t *testing.T) {
		page := 0
		limit := 0
		category := ""

		repo.On("SelectAll", 1, 8, category).Return(returnData, 0, nil).Once()

		result, _, err := srv.GetAll(page, limit, category)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		page := 1
		limit := 8
		category := "phones"

		repo.On("SelectAll", page, limit, category).Return(nil, 0,  errors.New("database error")).Once()

		result, _, err := srv.GetAll(page, limit, category)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		page := 1
		limit := 8
		category := "smartphones"

		repo.On("SelectAll", page, limit, category).Return(returnData, 0, nil).Once()

		result, _, err := srv.GetAll(page, limit, category)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})
}


func TestGetById(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := New(repo)

	returnData := &product.Core{
		ID:           1,
		Name:         "samsung",
		Description:  "hp ini murah sekali",
		Category:     "phones",
		Stock:        20,
		Price:        100000,
		PhotoProduct: "https://res.cloudinary.com/ikyft.jpg",
		User: user.Core{
			ID:           1,
			Name:         "lendra",
			PhotoProfile: "https://res.cloudinary.com/ikyft.jpg",
		},
	}

	t.Run("error from repository", func(t *testing.T) {
		idProduct := 1

		repo.On("SelectById", idProduct).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetById(idProduct)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		idProduct := 1

		repo.On("SelectById", idProduct).Return(returnData, nil).Once()

		result, err := srv.GetById(idProduct)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := New(repo)

	returnData := product.Core{
		ID:           1,
		Name:         "samsung",
		Description:  "hp ini murah sekali",
		Category:     "phones",
		Stock:        20,
		Price:        100000,
		PhotoProduct: "https://res.cloudinary.com/ikyft.jpg",
		UserID:       1,
	}

	t.Run("error from repository on update", func(t *testing.T) {
		repo.On("Update", 1, returnData).Return(errors.New("database error")).Once()

		err := srv.Update(1, returnData)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Update", 1, returnData).Return(nil).Once()

		err := srv.Update(1, returnData)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := New(repo)

	t.Run("error from repository on delete", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(errors.New("database error")).Once()

		err := srv.Delete(1, 1)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Delete", 1, 1).Return(nil).Once()

		err := srv.Delete(1, 1)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetByUserId(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := New(repo)

	returnData := []product.Core{
		{
			ID:           1,
			Name:         "samsung",
			Category:     "phones",
			Price:        100000,
			PhotoProduct: "https://res.cloudinary.com/ikyft.jpg",
		},
		{
			ID:           2,
			Name:         "vivo",
			Category:     "phones",
			Price:        120000,
			PhotoProduct: "https://res.cloudinary.com/zzsds.jpg",
		},
	}

	t.Run("error from repository", func(t *testing.T) {
		userIdLogin := 1

		repo.On("SelectByUserId", userIdLogin).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetByUserId(userIdLogin)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		userIdLogin := 1

		repo.On("SelectByUserId", userIdLogin).Return(returnData, nil).Once()

		result, err := srv.GetByUserId(userIdLogin)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})
}

func TestSearch(t *testing.T) {
	repo := new(mocks.ProductData)
	srv := New(repo)

	returnData := []product.Core{
		{
			ID:           1,
			Name:         "samsung",
			Category:     "phones",
			Price:        100000,
			PhotoProduct: "https://res.cloudinary.com/ikyft.jpg",
		},
		{
			ID:           2,
			Name:         "vivo",
			Category:     "phones",
			Price:        120000,
			PhotoProduct: "https://res.cloudinary.com/zzsds.jpg",
		},
	}

	t.Run("error from repository", func(t *testing.T) {
		query := "samsung"

		repo.On("Search", query).Return(nil, errors.New("database error")).Once()

		result, err := srv.Search(query)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		query := "samsung"

		repo.On("Search", query).Return(returnData, nil).Once()

		result, err := srv.Search(query)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})
}
