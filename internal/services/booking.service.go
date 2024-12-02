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

func (s *ServiceBooking) GetBookings(userid string) (*common.RespList, error) {
	bookings, err := s.BookingRepo.GetBookings(models.BookingWhere{
		UserID: userid,
	})
	if err != nil {
		return nil, err
	}

	resp := &common.RespList{
		TotalItem: len(bookings),
		Success:   false,
		Message:   "-",
		Data:      nil,
	}

	if len(bookings) == 0 {
		resp.Message = "Tidak ada data"
		return resp, nil
	}

	resp.TotalItem = len(bookings)
	resp.Success = true
	resp.Message = "Bookings exists"
	resp.Data = bookings

	return resp, nil
}

func (s *ServiceBooking) GetBooking(userid string, id string) (*common.RespDetail, error) {
	booking, err := s.BookingRepo.GetBookingDetails(models.BookingWhere{
		ID:     id,
		UserID: userid,
	})
	if err != nil {
		return nil, err
	}

	resp := &common.RespDetail{
		Success: false,
		Message: "-",
		Data:    nil,
	}

	if booking.ID == "" {
		resp.Message = "Tidak ada data"
		return resp, nil
	}

	resp.Success = true
	resp.Message = "Bookings exists"
	resp.Data = booking

	return resp, nil
}
func (s *ServiceBooking) UpdateBooking(userid string, id string, req *dtos.UpdateBookingReq) (*common.RespDetail, error) {

	bookingBefore, err := s.BookingRepo.GetBookingDetails(models.BookingWhere{ID: id})
	if err != nil {
		return nil, err
	}
	car, err := s.carsRepo.GetCar(bookingBefore.CarsID)
	if err != nil {
		return nil, err
	}

	booking := &models.Booking{
		CarsID:        req.CarsID,
		UserID:        req.UserID,
		Status:        req.Status,
		PaymentMethod: req.PaymentMethod,
	}

	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02 15:04:05", req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("time format error", err.Error())
		}

		totalNew := recalculateTotal(car.DailyRate, bookingBefore.EndDate, startDate)
		if bookingBefore.TotalPrice != totalNew {
			booking.TotalPrice = totalNew
		}
		booking.StartDate = startDate
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02 15:04:05", req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("time format error", err.Error())
		}

		totalNew := recalculateTotal(car.DailyRate, endDate, bookingBefore.StartDate)
		if bookingBefore.TotalPrice != totalNew {
			booking.TotalPrice = totalNew
		}
		booking.EndDate = endDate
	}

	err = s.BookingRepo.UpdateBooking(id, booking)
	if err != nil {
		return nil, err
	}

	resp := &common.RespDetail{
		Success: false,
		Message: "-",
		Data:    nil,
	}

	if booking.ID == "" {
		resp.Message = "Tidak ada data"
		return resp, nil
	}

	resp.Success = true
	resp.Message = "Updated"
	resp.Data = nil

	return resp, nil
}

func recalculateTotal(rate float64, end, start time.Time) float64 {
	dayElapse := math.Ceil(end.Sub(start).Hours() / 24)
	return rate * dayElapse
}
