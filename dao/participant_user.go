package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/YanshuoH/youkonger/models"
)

type participantUser struct {
	*gorm.DB
}

func (pu *participantUser) FindById(id uint64) (*models.ParticipantUser, error) {
	res := &models.ParticipantUser{}
	err := pu.First(res, id).Error
	return res, err
}

func (pu *participantUser) FindByUUID(uuid string) (*models.ParticipantUser, error) {
	res := &models.ParticipantUser{}
	err := pu.Where("uuid = ? ANd removed = FALSE", uuid).First(&res).Error
	return res, err
}

func (pu *participantUser) FindByUserUUIDAndEventUUID(uuid, eventUuid string) (*models.ParticipantUser, error) {
	res := &models.ParticipantUser{}
	err := pu.
		Joins("INNER JOIN event e ON e.id = participant_user.event_id").
		Joins("INNER JOIN user u ON u.id = participant_user.user_id").
		Where("u.uuid = ? AND e.uuid = ? AND participant_user.removed = FALSE",
			uuid, eventUuid).
		First(&res).Error
	return res, err
}

var ParticipantUser *participantUser

func initParticipantUser(conn *gorm.DB) {
	ParticipantUser = &participantUser{conn}
}
