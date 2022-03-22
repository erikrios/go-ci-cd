package service

type AuthService interface {
	GenerateToken(id uint, email string) (token string, err error)
}
