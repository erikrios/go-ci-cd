package service

import (
	"errors"
	"testing"

	"github.com/erikrios/go-clean-arhictecture/model"
	"github.com/erikrios/go-clean-arhictecture/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserServiceImpl(t *testing.T) {
	mockRepo := &mocks.UserRepository{}

	t.Run("success scenario", func(t *testing.T) {
		userService := NewUserServiceImpl(mockRepo)
		assert.NotNil(t, userService)
	})
}

func TestCreate(t *testing.T) {
	mockRepo := &mocks.UserRepository{}
	dummyRequest := model.CreateUserRequest{
		Email:    "naruto@gmail.com",
		Password: "naruto",
	}

	t.Run("success scenario", func(t *testing.T) {
		mockRepo.On("Insert", mock.AnythingOfType("*entities.User")).Return(nil).Once()

		service := NewUserServiceImpl(mockRepo)
		resp, err := service.Create(dummyRequest)

		assert.NoError(t, err)
		assert.Equal(t, dummyRequest.Email, resp.Email)
	})

	t.Run("failed scenario", func(t *testing.T) {
		mockRepo.On("Insert", mock.AnythingOfType("*entities.User")).Return(errors.New("something went wrong.")).Once()

		service := NewUserServiceImpl(mockRepo)
		_, err := service.Create(dummyRequest)

		assert.Error(t, err)
	})
}

func TestList(t *testing.T) {
	mockRepo := &mocks.UserRepository{}

	t.Run("success scenario", func(t *testing.T) {
		mockRepo.On("FindAll", mock.AnythingOfType("*[]entities.User")).Return(nil).Once()

		service := NewUserServiceImpl(mockRepo)
		_, err := service.List()

		assert.NoError(t, err)
	})

	t.Run("failed scenario", func(t *testing.T) {
		mockRepo.On("FindAll", mock.AnythingOfType("*[]entities.User")).Return(errors.New("something went wrong.")).Once()

		service := NewUserServiceImpl(mockRepo)
		_, err := service.List()

		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	mockRepo := &mocks.UserRepository{}
	dummyRequest := model.LoginUserRequest{
		Email:    "naruto@gmail.com",
		Password: "naruto",
	}

	t.Run("success scenario", func(t *testing.T) {
		mockRepo.On("Login", mock.AnythingOfType("*entities.User")).Return(nil).Once()

		service := NewUserServiceImpl(mockRepo)
		resp, err := service.Login(dummyRequest)

		assert.NoError(t, err)
		assert.Equal(t, dummyRequest.Email, resp.Email)
	})

	t.Run("failed scenario", func(t *testing.T) {
		mockRepo.On("Login", mock.AnythingOfType("*entities.User")).Return(errors.New("username and password not match.")).Once()

		service := NewUserServiceImpl(mockRepo)
		_, err := service.Login(dummyRequest)

		assert.Error(t, err)
	})
}
