package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelsWithID struct {
	ID        string         `json:"id" gorm:"primaryKey;varchar(40);"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:null;"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *ModelsWithID) generateID(prefix string) string {
	return prefix + "_" + uuid.New().String()
}

func (m *ModelsWithID) GenerateUUID(prefix string) {
	m.ID = m.generateID(prefix)
}
