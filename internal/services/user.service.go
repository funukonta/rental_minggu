package services

import (
	"errors"
	"rentalMobil/internal/dtos"
	"rentalMobil/internal/models"
	"rentalMobil/internal/repositories"
	"rentalMobil/pkg/auth"
	"rentalMobil/pkg/common"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ServiceUser struct {
	userRepo repositories.RepoUser
}

func NewServiceUser(repo repositories.RepoUser) *ServiceUser {
	return &ServiceUser{userRepo: repo}
}

func (s *ServiceUser) Login(req dtos.LoginReq) (*common.RespCreate, error) {
	user, err := s.userRepo.GetUser(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Username / Password Salah!")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("Username / Password Salah!")
	}

	authToken, err := auth.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	resp := &common.RespCreate{
		Success:     true,
		Message:     "Log in!",
		ReferenceID: user.ID,
		Data:        authToken,
	}

	return resp, nil

}

func (s *ServiceUser) Register(req dtos.RegisterReq) (*common.RespCreate, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.Users{
		Password: string(hashed),
		Email:    req.Email,
		Username: req.Username,
		Role:     req.Role,
	}
	err = s.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	resp := &common.RespCreate{
		Success:     true,
		Message:     "Created!",
		ReferenceID: user.ID,
		Data:        user,
	}

	return resp, nil
}
