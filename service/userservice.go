package service

import (
	"errors"
	"goexercise/models"

	"github.com/rs/zerolog/log"
)

type UserService struct {
	repo IUserRepository
}

func NewUserService(repo *IUserRepository) *UserService {
	return &UserService{
		repo: *repo,
	}
}

func (svc *UserService) GetUsers() *[]models.User {
	return svc.repo.GetUsers()
}

func (svc *UserService) GetUser(name string) (*models.User, error) {
	user, err := svc.repo.GetUser(name)

	if err != nil {
		log.Error().Err(err)
		return nil, errors.New("user name not exists")
	}
	return user, nil
}
