package services

import (
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type UserService interface {
	Create(UserInput web.CreateUserRequest) (models.User, error)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (s *UserServiceImpl) Create(UserInput web.CreateUserRequest) (models.User, error) {
	user := models.User{
		Username: UserInput.Username,
		Email:    UserInput.Email,
		Password: UserInput.Password,
		Age:      UserInput.Age,
	}

	newUser, err := s.UserRepository.Create(user)
	return newUser, err
}
