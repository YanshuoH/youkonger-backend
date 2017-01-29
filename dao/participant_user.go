package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/YanshuoH/youkonger/models"
)

type participantUser struct {
	*gorm.DB
}

func (pu *participantUser) FindByUUID(uuid string) (*models.ParticipantUser, error) {
	res := &models.ParticipantUser{}
	err := pu.Where("uuid = ? ANd removed = FALSE", uuid).First(&res).Error
	return res, err
}

var ParticipantUser *participantUser

func initParticipantUser(conn *gorm.DB) {
	ParticipantUser = &participantUser{conn}
}
