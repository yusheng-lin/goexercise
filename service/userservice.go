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

func (svc *UserService) GetUsers(page *models.Paging) (int64, int64, *[]models.User) {
	return svc.repo.GetUsers(page)
}

func (svc *UserService) GetUser(name string) (*models.User, error) {
	user, err := svc.repo.GetUser(name)

	if err != nil {
		log.Error().Err(err)
		return nil, errors.New("user name not exists")
	}
	return user, nil
}
