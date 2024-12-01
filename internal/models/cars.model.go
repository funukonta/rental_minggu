package models

import (
	"rentalMobil/pkg/common"

	"gorm.io/gorm"
)

type Cars struct {
	common.ModelsWithID
	Model      string  `json:"model"`
	PlatNumber string  `json:"plat_number" gorm:"unique"`
	DailyRate  float64 `json:"daily_rate"`
	Status     bool    `json:"status"`
	Color      string  `json:"color"`
	Make       string  `json:"make"`
	Year       int     `json:"year"`
}

func (c *Cars) TableName() string {
	return "Cars"
}

func (c *Cars) BeforeCreate(db *gorm.DB) error {
	c.GenerateUUID("cars")
	return nil
}
