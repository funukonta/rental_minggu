package repositories

import (
	"rentalMobil/internal/models"

	"gorm.io/gorm"
)

type RepoCars interface {
	Create(car *models.Cars) error
	GetCars() ([]models.Cars, error)
	GetCar(id string) (*models.Cars, error)
	UpdateCar(car models.Cars, id string) error
	DeleteCar(id string) error
}

type repoCars struct {
	dB *gorm.DB
}

func NewRepoCars(db *gorm.DB) RepoCars {
	return &repoCars{dB: db}
}

func (r *repoCars) Create(car *models.Cars) error {
	return r.dB.Create(car).Error
}

func (r *repoCars) GetCars() ([]models.Cars, error) {
	users := []models.Cars{}
	err := r.dB.Find(&users).Error

	return users, err
}

func (r *repoCars) GetCar(id string) (*models.Cars, error) {
	car := &models.Cars{}
	err := r.dB.Find(&car, "id = ?", id).Error

	return car, err
}

func (r *repoCars) UpdateCar(car models.Cars, id string) error {
	return r.dB.Updates(car).Error
}

func (r *repoCars) DeleteCar(id string) error {
	return r.dB.Where("id = ?", id).Delete(&models.Cars{}).Error
}
