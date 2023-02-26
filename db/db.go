package db

import (
	"goexercise/models"
	"goexercise/service"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *service.IUserRepository {
	var repo service.IUserRepository = &UserRepository{
		db: db,
	}
	return &repo
}

func (repo *UserRepository) GetUsers() *[]models.User {
	users := &[]models.User{}
	repo.db.Table("users").Find(users)
	return users
}

func (repo *UserRepository) GetUser(name string) (*models.User, error) {
	user := &models.User{}
	if err := repo.db.Table("users").Where("fullname=?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAccount(account string) (*models.Account, error) {
	acct := &models.Account{}
	if err := repo.db.Table("users").Where("acct=?", account).First(acct).Error; err != nil {
		return nil, err
	}
	return acct, nil
}

func (repo *UserRepository) CreateAccount(account *models.Account) error {
	err := repo.db.Table("users").Select("acct", "fullname", "pwd").Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteAccount(account string) error {
	err := repo.db.Table("users").Where("acct=?", account).Delete(&models.Account{}).Error
	if err != nil {
		return err
	}
	return nil
}
