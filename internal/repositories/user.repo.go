package repositories

import (
	"rentalMobil/internal/models"

	"gorm.io/gorm"
)

type RepoUser interface {
	Register(*models.Users) error
	GetUser(username string) (*models.Users, error)
}

type repoUser struct {
	dB *gorm.DB
}

func NewRepoUser(db *gorm.DB) RepoUser {
	return &repoUser{dB: db}
}

func (r *repoUser) Register(user *models.Users) error {
	err := r.dB.Create(user).Error
	return err
}

func (r *repoUser) GetUser(username string) (*models.Users, error) {
	user := &models.Users{}

	err := r.dB.Find(user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
