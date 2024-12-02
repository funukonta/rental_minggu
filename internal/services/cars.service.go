package services

import (
	"fmt"
	"rentalMobil/internal/dtos"
	"rentalMobil/internal/models"
	"rentalMobil/internal/repositories"
	"rentalMobil/pkg/common"

	"gorm.io/gorm"
)

type ServiceCars struct {
	carsRepo repositories.RepoCars
}

func NewServiceCars(repo repositories.RepoCars) *ServiceCars {
	return &ServiceCars{carsRepo: repo}
}

func (s *ServiceCars) CreateCars(req *dtos.CreateCarReq) (*common.RespCreate, error) {
	cars := &models.Cars{
		Model:      req.Model,
		PlatNumber: req.PlatNumber,
		DailyRate:  req.DailyRate,
		Status:     req.Status,
		Color:      req.Color,
		Make:       req.Make,
		Year:       req.Year,
	}
	err := s.carsRepo.Create(cars)
	if err != nil {
		return nil, err
	}

	resp := &common.RespCreate{
		Success:     true,
		Message:     "Cars Created!",
		ReferenceID: cars.ID,
		Data:        cars,
	}

	return resp, nil
}

func (s *ServiceCars) GetCars() (*common.RespList, error) {
	cars, err := s.carsRepo.GetCars()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("cars tidak ada")
		}
		return nil, err
	}

	resp := &common.RespList{
		Success:   true,
		Message:   "Cars Exists",
		TotalItem: len(cars),
		Data:      cars,
	}

	return resp, nil

}

func (s *ServiceCars) GetCar(id string) (*common.RespDetail, error) {
	car, err := s.carsRepo.GetCar(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("cars tidak ada")
		}
		return nil, err
	}

	if car.ID == "" {
		return nil, fmt.Errorf("cars tidak ada")
	}

	resp := &common.RespDetail{
		Success: true,
		Message: "Car ID exists",
		Data:    car,
	}

	return resp, nil
}

func (s *ServiceCars) UpdateCar(req *dtos.UpdateCarReq, id string) (*common.RespCreate, error) {
	cars := models.Cars{
		Model:     req.Model,
		DailyRate: req.DailyRate,
		Status:    req.Status,
		Color:     req.Color,
		Make:      req.Make,
		Year:      req.Year,
	}
	err := s.carsRepo.UpdateCar(&cars, id)
	if err != nil {
		return nil, err
	}

	resp := &common.RespCreate{
		Success:     true,
		Message:     "update ok",
		ReferenceID: id,
		Data:        cars,
	}

	return resp, nil
}

func (s *ServiceCars) DeleteCar(id string) (*common.RespCreate, error) {
	err := s.carsRepo.DeleteCar(id)
	if err != nil {
		return nil, err
	}

	resp := &common.RespCreate{
		Success:     true,
		Message:     "delete ok",
		ReferenceID: id,
		Data:        nil,
	}

	return resp, nil
}
