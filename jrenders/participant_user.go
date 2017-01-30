package jrenders

import (
	"github.com/YanshuoH/youkonger/dao"
	"github.com/YanshuoH/youkonger/models"
)

type participantUser struct{}

type JParticipantUser struct {
	UUID     string `json:"participantUserUuid"`
	UserUUID string `json:"userUuid"`
	Name     string `json:"name"`
}

func (r *participantUser) Itemize(pu *models.ParticipantUser) JParticipantUser {
	if pu.User == nil {
		u, _ := dao.User.FindById(pu.UserId)
		pu.User = u
	}

	return JParticipantUser{
		UUID:     pu.UUID,
		UserUUID: pu.User.UUID,
		Name:     pu.Name,
	}
}

var ParticipantUser *participantUser

func initParticipantUser() {
	ParticipantUser = &participantUser{}
}
