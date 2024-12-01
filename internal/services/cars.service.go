package services

import (
	"rentalMobil/internal/dtos"
	"rentalMobil/internal/models"
	"rentalMobil/internal/repositories"
	"rentalMobil/pkg/common"
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
