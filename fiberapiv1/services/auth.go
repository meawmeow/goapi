package services

import (
	"errors"
	"fiberapiv1/entity"
	"fiberapiv1/errs"
	"fiberapiv1/logs"
	"fiberapiv1/models"
	"fiberapiv1/repository"
	"fiberapiv1/security"
	"strconv"
)

type UserService interface {
	CreateUser(models.UserRequest) (*models.UserResponse, error)
	Login(models.UserRequest) (*models.UserResponse, error)
	GetUsers() ([]models.UserResponse, error)
	GetUser(int) (*models.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(request models.UserRequest) (*models.UserResponse, error) {

	password, _ := security.EncryptPassword(request.Password)
	userMaping := entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: []byte(password),
		Address:  request.Address,
	}

	user, err := s.userRepo.InsertUser(userMaping)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := models.UserResponse{
		Id:       int(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
	}
	return &response, nil
}

func (s userService) Login(request models.UserRequest) (*models.UserResponse, error) {

	//step 1 get email
	user, err := s.userRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if user.Email == "" {
		return nil, errors.New("user not found")
	}
	//step 2 check password
	if err := security.VerifyPassword(string(user.Password), request.Password); err != nil {
		return nil, errors.New("incorrect password")
	}

	//step 3 create token
	token, err := security.GenerateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		return nil, err
	}

	//step final make response
	userResponse := models.UserResponse{
		Id:       int(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
		Token:    *token,
	}

	return &userResponse, nil
}

func (s userService) GetUsers() ([]models.UserResponse, error) {

	users, err := s.userRepo.GetAllUser()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	userResponses := []models.UserResponse{}
	for _, user := range users {
		userResponse := models.UserResponse{
			Id:       int(user.ID),
			Username: user.Username,
			Email:    user.Email,
			Address:  user.Address,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (s userService) GetUser(id int) (*models.UserResponse, error) {

	user, err := s.userRepo.UserByID(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	userResponse := models.UserResponse{
		Id:       int(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
	}
	return &userResponse, nil
}
