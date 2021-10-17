// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	bookauthor "github.com/daniel5u/suisei/domain/bookauthor"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteByBookID provides a mock function with given fields: bookid
func (_m *Repository) DeleteByBookID(bookid int) error {
	ret := _m.Called(bookid)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(bookid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: bookauthorDomain
func (_m *Repository) Store(bookauthorDomain bookauthor.Domain) (bookauthor.Domain, error) {
	ret := _m.Called(bookauthorDomain)

	var r0 bookauthor.Domain
	if rf, ok := ret.Get(0).(func(bookauthor.Domain) bookauthor.Domain); ok {
		r0 = rf(bookauthorDomain)
	} else {
		r0 = ret.Get(0).(bookauthor.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bookauthor.Domain) error); ok {
		r1 = rf(bookauthorDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}