package service

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/erikrios/go-clean-arhictecture/constants"
)

type AuthServiceImpl struct{}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (a *AuthServiceImpl) GenerateToken(id uint, email string) (token string, err error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	jwtWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtWithClaims.SignedString([]byte(constants.JWT_SECRET))
	return
}

// func (a *AuthServiceImpl) ExtractTokenUserId(c echo.Context) (id uint) {
// 	user := c.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		id = uint(claims["id"].(float64))
// 	}
// 	return
// }
