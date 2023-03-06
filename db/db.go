package db

import (
	"fmt"
	"goexercise/models"
	"goexercise/service"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) service.IUserRepository {
	var repo service.IUserRepository = &UserRepository{
		db: db,
	}
	return repo
}

func (repo *UserRepository) GetUsers(page *models.Paging) (int64, int64, []*models.User) {
	users := []*models.User{}
	offset := (page.PageNo - 1) * page.Take
	var rows int64
	repo.db.Table("users").Count(&rows)
	pages := rows / int64(page.Take)

	if rows%int64(page.Take) > 0 {
		pages += 1
	}

	repo.db.Table("users").
		Limit(page.Take).
		Offset(offset).
		Order(fmt.Sprintf(`%s %s`, page.SortBy, page.Sort)).Find(&users)

	return rows, pages, users
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

func (repo *UserRepository) UpdateAccount(account *models.Account) error {
	val := map[string]interface{}{"updated_at": time.Now()}

	if account.Fullname != "" {
		val["fullname"] = account.Fullname
	}

	if account.Pwd != "" {
		val["pwd"] = account.Pwd
	}

	err := repo.db.Table("users").Where("acct=?", account.Acct).Updates(val).Error
	if err != nil {
		return err
	}
	return nil
}
