package service

import (
	"errors"
	"goexercise/models"
	"goexercise/utility"

	"github.com/rs/zerolog/log"
)

type AccountService struct {
	repo IUserRepository
}

func NewAccountService(repo *IUserRepository) *AccountService {
	return &AccountService{
		repo: *repo,
	}
}

func (svc *AccountService) Login(account string, pwd string) (*models.Account, error) {
	acct, err := svc.repo.GetAccount(account)

	if err != nil {
		log.Error().Err(err)
		return nil, errors.New("account not exist")
	}

	hash := utility.GetHash(pwd)

	if acct.Pwd != hash {
		return nil, errors.New("password incorrect")
	}
	return acct, nil
}

func (svc *AccountService) SignUp(account *models.Account) error {
	acct, _ := svc.repo.GetAccount(account.Acct)

	if acct != nil {
		return errors.New("this account is already exist, please use a different name")
	}

	account.Pwd = utility.GetHash(account.Pwd)

	err := svc.repo.CreateAccount(account)

	if err != nil {
		log.Error().Err(err)
		return errors.New("create account fail. please SignUp later")
	}

	return nil
}
