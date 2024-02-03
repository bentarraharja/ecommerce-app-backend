// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	product "MyEcommerce/features/product"

	mock "github.com/stretchr/testify/mock"
)

// ProductData is an autogenerated mock type for the ProductDataInterface type
type ProductData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userIdLogin, IdProduct
func (_m *ProductData) Delete(userIdLogin int, IdProduct int) error {
	ret := _m.Called(userIdLogin, IdProduct)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(userIdLogin, IdProduct)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: userIdLogin, input
func (_m *ProductData) Insert(userIdLogin int, input product.Core) error {
	ret := _m.Called(userIdLogin, input)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, product.Core) error); ok {
		r0 = rf(userIdLogin, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: query
func (_m *ProductData) Search(query string) ([]product.Core, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 []product.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]product.Core, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) []product.Core); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAll provides a mock function with given fields: page, limit, category
func (_m *ProductData) SelectAll(page int, limit int, category string) ([]product.Core, int, error) {
	ret := _m.Called(page, limit, category)

	if len(ret) == 0 {
		panic("no return value specified for SelectAll")
	}

	var r0 []product.Core
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]product.Core, int, error)); ok {
		return rf(page, limit, category)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []product.Core); ok {
		r0 = rf(page, limit, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int); ok {
		r1 = rf(page, limit, category)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, limit, category)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SelectById provides a mock function with given fields: IdProduct
func (_m *ProductData) SelectById(IdProduct int) (*product.Core, error) {
	ret := _m.Called(IdProduct)

	if len(ret) == 0 {
		panic("no return value specified for SelectById")
	}

	var r0 *product.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*product.Core, error)); ok {
		return rf(IdProduct)
	}
	if rf, ok := ret.Get(0).(func(int) *product.Core); ok {
		r0 = rf(IdProduct)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(IdProduct)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByUserId provides a mock function with given fields: userIdLogin
func (_m *ProductData) SelectByUserId(userIdLogin int) ([]product.Core, error) {
	ret := _m.Called(userIdLogin)

	if len(ret) == 0 {
		panic("no return value specified for SelectByUserId")
	}

	var r0 []product.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]product.Core, error)); ok {
		return rf(userIdLogin)
	}
	if rf, ok := ret.Get(0).(func(int) []product.Core); ok {
		r0 = rf(userIdLogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userIdLogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userIdLogin, input
func (_m *ProductData) Update(userIdLogin int, input product.Core) error {
	ret := _m.Called(userIdLogin, input)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, product.Core) error); ok {
		r0 = rf(userIdLogin, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductData creates a new instance of ProductData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductData(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductData {
	mock := &ProductData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
