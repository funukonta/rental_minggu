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
	db.Preload("Cars")
	db.Preload("Users")
	return db.Error
}
