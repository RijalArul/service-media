package services

import (
	models "service-media/models/entity"
	"service-media/models/web"
	"service-media/repositories"
)

type UserService interface {
	Create(UserInput web.CreateUserRequest) (web.CreateUserResponse, error)
	Login(LoginInput web.LoginUserRequest) (models.User, error)
	GetUser(id uint) (web.CreateUserResponse, error)
	UpdateUser(UpdateInput web.UpdateUserRequest, userID uint) (web.UpdateUserResponse, error)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

const (
	appJSON = "appplication/json"
)

func convertUserBodyResponse(user models.User) web.CreateUserResponse {
	return web.CreateUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
}

func (s *UserServiceImpl) Create(UserInput web.CreateUserRequest) (web.CreateUserResponse, error) {
	user := models.User{
		Username: UserInput.Username,
		Email:    UserInput.Email,
		Password: UserInput.Password,
		Age:      UserInput.Age,
	}

	newUser, err := s.UserRepository.Create(user)

	userResp := web.CreateUserResponse{
		Id:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Age:      newUser.Age,
	}

	return userResp, err
}

func (s *UserServiceImpl) Login(LoginInput web.LoginUserRequest) (models.User, error) {
	user, err := s.UserRepository.FindByEmail(LoginInput.Email)
	return user, err
}

func (s *UserServiceImpl) GetUser(id uint) (web.CreateUserResponse, error) {
	user, err := s.UserRepository.FindByID(id)
	userResp := web.CreateUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
	return userResp, err
}

func (s *UserServiceImpl) UpdateUser(UpdateInput web.UpdateUserRequest, userID uint) (web.UpdateUserResponse, error) {
	newUpdateUser := models.User{
		Username: UpdateInput.Username,
		Email:    UpdateInput.Email,
	}

	updateUser, err := s.UserRepository.UpdateByID(newUpdateUser, userID)

	userResp := web.UpdateUserResponse{
		Id:        updateUser.ID,
		Username:  updateUser.Username,
		Email:     updateUser.Email,
		Age:       updateUser.Age,
		UpdatedAt: *updateUser.UpdatedAt,
	}
	return userResp, err
}
