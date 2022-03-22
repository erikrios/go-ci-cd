package service

import (
	"github.com/erikrios/go-clean-arhictecture/entities"
	"github.com/erikrios/go-clean-arhictecture/model"
	"github.com/erikrios/go-clean-arhictecture/repository"
)

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) *userServiceImpl {
	return &userServiceImpl{repository: repository}
}

func (u *userServiceImpl) Create(request model.CreateUserRequest) (response model.GetUserResponse, err error) {
	user := &entities.User{
		Email:    request.Email,
		Password: request.Password,
	}

	err = u.repository.Insert(user)

	response.Id = user.ID
	response.Email = user.Email
	response.CreatedAt = user.CreatedAt
	response.UpdatedAt = user.UpdatedAt

	return
}

func (u *userServiceImpl) List() (responses []model.GetUserResponse, err error) {
	responses = make([]model.GetUserResponse, 0)
	users := make([]entities.User, 0)

	err = u.repository.FindAll(&users)

	for _, user := range users {
		response := model.GetUserResponse{
			Id:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return
}

func (u *userServiceImpl) Login(request model.LoginUserRequest) (user model.GetUserResponse, err error) {
	userEntity := &entities.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err = u.repository.Login(userEntity); err != nil {
		return
	}

	user.Id = userEntity.ID
	user.Email = userEntity.Email
	return
}
