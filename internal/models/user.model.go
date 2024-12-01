package models

import (
	"rentalMobil/pkg/common"
	"time"

	"gorm.io/gorm"
)

const (
	ROLE_ADMIN = "ADMIN"
	ROLE_CUST  = "CUSTOMER"
)

type Users struct {
	common.ModelsWithID
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	Email     string    `json:"email" gorm:"unique"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	LastLogin time.Time `json:"last_login" gorm:"default:null;"`
}

func (m *Users) BeforeCreate(db *gorm.DB) error {
	m.GenerateUUID("user")
	return nil
}
