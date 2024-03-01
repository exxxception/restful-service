package services

import (
	"github.com/exxxception/restful-service/models"
	"github.com/exxxception/restful-service/pkg/repository"
)

type Accounts interface {
	Create(amount float64) (int, error)
	GetInfo(accountId int) (models.Account, error)
	Deposit(accountId int, amount float64) error
	Withdraw(accountId int, amount float64) error
}

type Services struct {
	Accounts
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		Accounts: NewAccountsService(repos.Accounts),
	}
}
