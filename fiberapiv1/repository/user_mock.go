package repository

import (
	"fiberapiv1/entity"

	"gorm.io/gorm"
)

type UserRepositoryMock interface {
	InsertUser(entity.User) (*entity.User, error)
	GetAllUser() ([]entity.User, error)
	UserByID(int) (*entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type userRepositoryMock struct {
	db *gorm.DB
}

func NewUserRepositoryMock(db *gorm.DB) UserRepositoryMock {
	return userRepositoryMock{db}
}

func (r userRepositoryMock) InsertUser(user entity.User) (*entity.User, error) {

	return &user, nil
}

func (r userRepositoryMock) GetAllUser() ([]entity.User, error) {

	user := []entity.User{
		{ID: 1, Username: "name 1", Email: "name1@gmail.com", Address: "Address1"},
		{ID: 2, Username: "name 2", Email: "name2@gmail.com", Address: "Address2"},
		{ID: 3, Username: "name 3", Email: "name3@gmail.com", Address: "Address3"},
	}
	return user, nil
}

func (r userRepositoryMock) UserByID(id int) (*entity.User, error) {
	user := entity.User{ID: uint64(id), Username: "name 1", Email: "name1@gmail.com", Address: "Address1"}
	return &user, nil
}

func (r userRepositoryMock) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if email == "name1@gmail.com" {
		user = entity.User{ID: 1, Username: "name 1", Email: "name1@gmail.com", Address: "Address1", Password: []byte("aaaa")}
	}
	return user, nil
}
