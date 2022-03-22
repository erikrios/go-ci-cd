// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// ExtractTokenUserId provides a mock function with given fields: c
func (_m *AuthService) ExtractTokenUserId(c echo.Context) uint {
	ret := _m.Called(c)

	var r0 uint
	if rf, ok := ret.Get(0).(func(echo.Context) uint); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// GenerateToken provides a mock function with given fields: id, email
func (_m *AuthService) GenerateToken(id uint, email string) (string, error) {
	ret := _m.Called(id, email)

	var r0 string
	if rf, ok := ret.Get(0).(func(uint, string) string); ok {
		r0 = rf(id, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(id, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
