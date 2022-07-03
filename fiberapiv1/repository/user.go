package repository

import (
	"fiberapiv1/entity"
	"fiberapiv1/logs"

	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(entity.User) (user *entity.User, err error)
	GetAllUser() (users []entity.User, err error)
	UserByID(int) (user *entity.User, err error)
	FindByEmail(email string) (user entity.User, err error)
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db}
}

func (r userRepositoryDB) InsertUser(enuser entity.User) (user *entity.User, err error) {
	logs.Info("Save init : " + enuser.Username)
	err = r.db.Save(&enuser).Error
	if err != nil {
		logs.Info("Save err : " + err.Error())
	}
	return &enuser, err
}

func (r userRepositoryDB) GetAllUser() (users []entity.User, err error) {
	err = r.db.Find(&users).Error
	return users, err
}

func (r userRepositoryDB) UserByID(id int) (user *entity.User, err error) {
	err = r.db.Where("id=?", id).First(&user).Error
	return user, err
}

func (r userRepositoryDB) FindByEmail(email string) (user entity.User, err error) {
	err = r.db.Where("email=?", email).First(&user).Error
	return user, err
}
