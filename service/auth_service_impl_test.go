package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthServiceImpl(t *testing.T) {
	t.Run("success scenario", func(t *testing.T) {
		authService := NewAuthServiceImpl()
		assert.NotNil(t, authService)
	})
}

func TestGenerateToken(t *testing.T) {
	t.Run("success scenario", func(t *testing.T) {
		service := NewAuthServiceImpl()
		token, err := service.GenerateToken(1, "naruto@gmail.com")

		assert.NoError(t, err)
		assert.NotZero(t, token)
	})
}
