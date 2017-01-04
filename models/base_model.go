package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID      string `gorm:"unique_index"`
	Removed   bool
}

func (m *BaseModel) BeforeCreate() {
	m.UUID = uuid.NewV4().String()

	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now
}

func (m *BaseModel) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
