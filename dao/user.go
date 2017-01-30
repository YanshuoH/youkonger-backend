package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/YanshuoH/youkonger/models"
)

type user struct {
	*gorm.DB
}

func (u *user) FindById(id uint64) (*models.User, error) {
	res := &models.User{}
	err := u.First(res, id).Error
	return res, err
}

func (u *user) FindByUUID(uuid string) (*models.User, error) {
	res := &models.User{}
	err := u.Where("uuid = ? AND removed = FALSE", uuid).First(res).Error

	return res, err
}

var User *user

func initUser(conn *gorm.DB) {
	User = &user{conn}
}
