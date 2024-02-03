package service

import (
	"MyEcommerce/features/user"
	"MyEcommerce/mocks"
	hashMock "MyEcommerce/utils/encrypts/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(hashMock.HashMock)

	returnData := user.Core{
		ID:           1,
		Name:         "arja",
		UserName:     "arja05",
		Email:        "arja@gmail.com",
		Password:     "qwerty",
		Role:         "user",
		PhotoProfile: "https://res.cloudinary.com/dlxvvuhph/image/upload/v1706070645/BE20_MyEcommerce/ebpux5x5e0uycmttm3g4.png",
	}

	t.Run("Success Get By Id", func(t *testing.T) {
		repo.On("SelectById", 1).Return(&returnData, nil).Once()
		srv := New(repo, hash)
		result, err := srv.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, returnData.Name, result.Name)
		assert.Equal(t, returnData.Email, result.Email)
		repo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		repo.On("SelectById", 1).Return(nil, errors.New("user not found")).Once()
		srv := New(repo, hash)
		result, err := srv.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(hashMock.HashMock)
	srv := New(repo, hash)

	inputData := user.Core{
		Name:     "arja",
		UserName: "arja05",
		Email:    "arja@gmail.com",
		Password: "password123",
		Role:     "user",
	}

	t.Run("Success Create User", func(t *testing.T) {
		hash.On("HashPassword", mock.AnythingOfType("string")).Return("hashedPassword", nil).Once()
		repo.On("Insert", mock.Anything).Return(nil).Once()
		inputData.Role = ""
		err := srv.Create(inputData)
		inputData.Role = "user"

		assert.NoError(t, err)
		hash.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("Validation Error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(nil).Once()
		invalidInput := user.Core{}
		err := srv.Create(invalidInput)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required")
		repo.AssertNotCalled(t, "Insert")
	})

	t.Run("Hash Password Error", func(t *testing.T) {
		hash.On("HashPassword", mock.Anything).Return("", errors.New("hash error")).Once()
		err := srv.Create(inputData)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Error hash password.")
		repo.AssertNotCalled(t, "Insert")
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(hashMock.HashMock)
	userService := New(repo, hash)

	input := user.Core{
		ID:           1,
		Name:         "UpdatedName",
		UserName:     "updatedUsername",
		Email:        "updated@gmail.com",
		Password:     "newpassword",
		Role:         "user",
		PhotoProfile: "https://example.com/newimage.png",
	}

	t.Run("invalid user id", func(t *testing.T) {
		err := userService.Update(0, input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid id")
	})

	t.Run("hash password error", func(t *testing.T) {
		hash.On("HashPassword", input.Password).Return("", errors.New("hash error")).Once()

		err := userService.Update(1, input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Error hash password.")
	})

	t.Run("error from repository", func(t *testing.T) {
		hashedPassword := "hashedPassword"
		hash.On("HashPassword", input.Password).Return(hashedPassword, nil).Once()
		repo.On("Update", 1, user.Core{ID: input.ID, Name: input.Name, UserName: input.UserName, Email: input.Email, Password: hashedPassword, Role: input.Role, PhotoProfile: input.PhotoProfile}).Return(errors.New("database error")).Once()

		err := userService.Update(1, input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
	})

	t.Run("success", func(t *testing.T) {
		hashedPassword := "hashedPassword"
		hash.On("HashPassword", input.Password).Return(hashedPassword, nil).Once()
		repo.On("Update", 1, user.Core{ID: input.ID, Name: input.Name, UserName: input.UserName, Email: input.Email, Password: hashedPassword, Role: input.Role, PhotoProfile: input.PhotoProfile}).Return(nil).Once()

		err := userService.Update(1, input)

		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(hashMock.HashMock)
	userService := New(repo, hash)

	t.Run("invalid user id", func(t *testing.T) {
		err := userService.Delete(0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid id")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Delete", 1).Return(nil).Once()

		err := userService.Delete(1)

		assert.NoError(t, err)
	})
}

func TestLogin(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(hashMock.HashMock)
	userService := New(repo, hash)

	inputLogin := user.Core{
		ID:       1,
		Email:    "updated@gmail.com",
		Password: "newpassword",
	}

	t.Run("empty email and password", func(t *testing.T) {
		_, _, err := userService.Login("", "")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email dan password wajib diisi")
	})

	t.Run("empty email", func(t *testing.T) {
		_, _, err := userService.Login("", "password")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email wajib diisi")
	})

	t.Run("empty password", func(t *testing.T) {
		_, _, err := userService.Login("email@gmail.com", "")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password wajib diisi")
	})

	t.Run("password not match", func(t *testing.T) {
		repo.On("Login", inputLogin.Email, inputLogin.Password).Return(&inputLogin, nil).Once()
		hash.On("CheckPasswordHash", inputLogin.Password, inputLogin.Password).Return(false).Once()

		_, _, err := userService.Login(inputLogin.Email, inputLogin.Password)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password tidak sesuai.")
	})

	t.Run("error on userData.Login", func(t *testing.T) {
		repo.On("Login", inputLogin.Email, inputLogin.Password).Return(nil, errors.New("some error")).Once()

		_, _, err := userService.Login(inputLogin.Email, inputLogin.Password)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "some error")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Login", inputLogin.Email, inputLogin.Password).Return(&inputLogin, nil).Once()
		hash.On("CheckPasswordHash", inputLogin.Password, inputLogin.Password).Return(true).Once()

		_, _, err := userService.Login(inputLogin.Email, inputLogin.Password)

		assert.NoError(t, err)
	})
}

func TestGetAdminUsers(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(hashMock.HashMock)
	userService := New(repo, hash)

	returnData := []user.Core{
		{
			ID:   1,
			Name: "manusia",
			Role: "user",
		},
		{
			ID:   2,
			Name: "manusia",
			Role: "user",
		},
	}

	userData := user.Core{
		ID:   1,
		Name: "admin",
		Role: "admin",
	}

	t.Run("default page and limit", func(t *testing.T) {
		page := 0
		limit := 0
		totalPage := 2

		repo.On("SelectById", 1).Return(&userData, nil).Once()
		repo.On("SelectAdminUsers", 1, 10).Return(returnData, nil, totalPage).Once()

		result, err, _ := userService.GetAdminUsers(1, page, limit)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})

	t.Run("error from SelectById", func(t *testing.T) {
		page := 1
		limit := 10

		repo.On("SelectById", 1).Return(nil, errors.New("user not found")).Once()

		result, err, _ := userService.GetAdminUsers(1, page, limit)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user not found", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("user role does not have access", func(t *testing.T) {
		page := 1
		limit := 10

		userData := user.Core{
			ID:   1,
			Name: "user",
			Role: "user",
		}

		repo.On("SelectById", 1).Return(&userData, nil).Once()

		result, err, _ := userService.GetAdminUsers(1, page, limit)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, "Sorry, your role does not have this access.", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		page := 1
		limit := 10
		totalPage := 2

		repo.On("SelectById", 1).Return(&userData, nil).Once()
		repo.On("SelectAdminUsers", page, limit).Return(nil, errors.New("database error"), totalPage).Once()

		result, err, _ := userService.GetAdminUsers(1, page, limit)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		page := 1
		limit := 10
		totalPage := 2

		repo.On("SelectById", 1).Return(&userData, nil).Once()
		repo.On("SelectAdminUsers", page, limit).Return(returnData, nil, totalPage).Once()

		result, err, _ := userService.GetAdminUsers(1, page, limit)

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)

		repo.AssertExpectations(t)
	})
}
