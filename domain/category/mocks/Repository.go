// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	category "github.com/daniel5u/suisei/domain/category"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields:
func (_m *Repository) Fetch() ([]category.Domain, error) {
	ret := _m.Called()

	var r0 []category.Domain
	if rf, ok := ret.Get(0).(func() []category.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]category.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id int) (category.Domain, error) {
	ret := _m.Called(id)

	var r0 category.Domain
	if rf, ok := ret.Get(0).(func(int) category.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(category.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *Repository) GetByName(name string) (category.Domain, error) {
	ret := _m.Called(name)

	var r0 category.Domain
	if rf, ok := ret.Get(0).(func(string) category.Domain); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(category.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: categoryDomain
func (_m *Repository) Store(categoryDomain category.Domain) (category.Domain, error) {
	ret := _m.Called(categoryDomain)

	var r0 category.Domain
	if rf, ok := ret.Get(0).(func(category.Domain) category.Domain); ok {
		r0 = rf(categoryDomain)
	} else {
		r0 = ret.Get(0).(category.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(category.Domain) error); ok {
		r1 = rf(categoryDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
