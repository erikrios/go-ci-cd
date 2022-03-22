package service

import "github.com/erikrios/go-clean-arhictecture/model"

type UserService interface {
	Create(request model.CreateUserRequest) (response model.GetUserResponse, err error)
	List() (responses []model.GetUserResponse, err error)
	Login(request model.LoginUserRequest) (user model.GetUserResponse, err error)
}
