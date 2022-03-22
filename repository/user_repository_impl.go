package repository

import (
	"github.com/erikrios/go-clean-arhictecture/entities"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (u *userRepositoryImpl) Insert(user *entities.User) error {
	return u.db.Create(user).Error
}

func (u *userRepositoryImpl) FindAll(users *[]entities.User) error {
	return u.db.Find(users).Error
}

func (u *userRepositoryImpl) Login(user *entities.User) error {
	return u.db.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error
}
