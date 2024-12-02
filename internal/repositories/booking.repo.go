package repositories

import (
	"rentalMobil/internal/models"

	"gorm.io/gorm"
)

type RepoBooking interface {
	Create(car *models.Booking) error
}

type repoBooking struct {
	dB *gorm.DB
}

func NewRepoBooking(db *gorm.DB) RepoBooking {
	return &repoBooking{dB: db}
}

func (r *repoBooking) Create(car *models.Booking) error {
	db := r.dB.Create(&car)
	db.Preload("Car") // Note : samakan dengan field di Struct Booking, di Booking field ini Car bukan Cars
	db.Preload("User").Find(&car)
	return db.Error
}
