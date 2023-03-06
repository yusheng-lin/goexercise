package service

import "goexercise/models"

type IUserRepository interface {
	GetUsers(*models.Paging) (int64, int64, []*models.User)
	GetUser(name string) (*models.User, error)
	GetAccount(account string) (*models.Account, error)
	CreateAccount(account *models.Account) error
	DeleteAccount(account string) error
	UpdateAccount(account *models.Account) error
}
