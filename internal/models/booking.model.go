package models

import (
	"rentalMobil/pkg/common"
	"time"

	"gorm.io/gorm"
)

const (
	STATUS_BOOKING_NOT_PAID = "NOT_PAID"
	STATUS_BOOKING_PAID     = "PAID"
	STATUS_BOOKING_DONE     = "DONE"
	STATUS_BOOKING_CANCEL   = "CALCEL"
)

type Booking struct {
	common.ModelsWithID
	CarsID        string    `json:"cars_id"`
	UserID        string    `json:"user_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	TotalPrice    float64   `json:"total_price"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`

	Cars Cars  `json:"car" gorm:"foreignKey:ID"`
	User Users `json:"user" gorm:"foreignKey:ID"`
}

func (c *Booking) TableName() string {
	return "booking"
}

func (c *Booking) BeforeCreate(db *gorm.DB) error {
	c.GenerateUUID("ob")
	return nil
}
