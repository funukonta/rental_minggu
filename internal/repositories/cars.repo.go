package repositories

import (
	"fmt"
	"rentalMobil/internal/models"

	"gorm.io/gorm"
)

type RepoCars interface {
	Create(car *models.Cars) error
	GetCars() ([]models.Cars, error)
	GetCar(id string) (*models.Cars, error)
	UpdateCar(car *models.Cars, id string) error
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
	sql := r.dB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&car, "id = ?", id)
	})

	fmt.Println(sql)

	return car, err
}

func (r *repoCars) UpdateCar(car *models.Cars, id string) error {
	return r.dB.Where("id = ?", id).Updates(car).Find(&car, "id = ?", id).Error
}

func (r *repoCars) DeleteCar(id string) error {
	// return r.dB.Where("id = ?", id).Delete(&models.Cars{}).Error
	tx := r.dB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("id=?", id).Delete(&models.Cars{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
