package repository

import "github.com/erikrios/go-clean-arhictecture/entities"

type UserRepository interface {
	Insert(user *entities.User) error
	FindAll(users *[]entities.User) error
	Login(user *entities.User) error
}
