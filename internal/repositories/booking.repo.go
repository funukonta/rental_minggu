package repositories

import (
	"rentalMobil/internal/models"

	"gorm.io/gorm"
)

type RepoBooking interface {
	Create(booking *models.Booking) error
	GetBookingDetails(where models.BookingWhere) (*models.Booking, error)
	GetBookings(where models.BookingWhere) ([]models.Booking, error)
	UpdateBooking(id string, booking *models.Booking) error
}

type repoBooking struct {
	dB *gorm.DB
}

func NewRepoBooking(db *gorm.DB) RepoBooking {
	return &repoBooking{dB: db}
}

func (r *repoBooking) Create(booking *models.Booking) error {
	db := r.dB.Create(&booking)
	populateQueryWithPreload(db)
	db.Find(&booking)

	return db.Error
}

func (r *repoBooking) GetBookingDetails(where models.BookingWhere) (*models.Booking, error) {
	booking := &models.Booking{}

	db := r.dB.Model(booking)
	populateQueryBookingwithWhere(db, where)

	populateQueryWithPreload(db)
	err := db.First(&booking).Error

	return booking, err
}

func (r *repoBooking) GetBookings(where models.BookingWhere) ([]models.Booking, error) {
	bookings := []models.Booking{} // booking doenst need preload yet, preload used for details, for performance purposes
	db := r.dB.Find(&bookings)
	populateQueryBookingwithWhere(db, where)
	err := db.Order("updated_at DESC").Error

	return bookings, err
}

func (r *repoBooking) UpdateBooking(id string, booking *models.Booking) error {
	db := r.dB.Model(booking).Where("id = ?", id)
	db.Updates(&booking)
	err := db.Error

	return err
}

func populateQueryBookingwithWhere(db *gorm.DB, where models.BookingWhere) {
	if where.ID != "" {
		db.Where("id = ?", where.ID)
	}

	if where.UserID != "" {
		db.Where("user_id = ?", where.UserID)
	}
}

func populateQueryWithPreload(db *gorm.DB) {
	db = db.Preload("Car") // Note : samakan dengan field di Struct Booking, di Booking field ini Car bukan Cars
	db = db.Preload("User")
}
