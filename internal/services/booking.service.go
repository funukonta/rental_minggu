package services

import (
	"fmt"
	"math"
	"rentalMobil/internal/dtos"
	"rentalMobil/internal/models"
	"rentalMobil/internal/repositories"
	"rentalMobil/pkg/common"
	"time"
)

type ServiceBooking struct {
	BookingRepo repositories.RepoBooking
	carsRepo    repositories.RepoCars
}

func NewServiceBooking(repo repositories.RepoBooking, carRepo repositories.RepoCars) *ServiceBooking {
	return &ServiceBooking{BookingRepo: repo, carsRepo: carRepo}
}

func (s *ServiceBooking) CreateBooking(req *dtos.CreateBookingReq) (*common.RespCreate, error) {

	startDate, err := time.Parse("2006-01-02 15:04:05", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("time format error", err.Error())
	}
	endDate, err := time.Parse("2006-01-02 15:04:05", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("time format error", err.Error())
	}

	cars, err := s.carsRepo.GetCar(req.CarsID)
	if err != nil {
		return nil, err
	}

	if cars.ID == "" {
		return nil, fmt.Errorf("car not exists")
	}

	dayElapse := math.Ceil(endDate.Sub(startDate).Hours() / 24)

	booking := &models.Booking{
		CarsID:        req.CarsID,
		UserID:        req.UserID,
		StartDate:     startDate,
		EndDate:       endDate,
		TotalPrice:    cars.DailyRate * dayElapse,
		PaymentMethod: req.PaymentMethod,
		Status:        models.STATUS_BOOKING_NOT_PAID,
	}
	err = s.BookingRepo.Create(booking)
	if err != nil {
		return nil, err
	}

	resp := &common.RespCreate{
		Success:     true,
		Message:     "Booking Created!",
		ReferenceID: booking.ID,
		Data:        booking,
	}

	return resp, nil
}
