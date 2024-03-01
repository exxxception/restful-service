package services

import (
	"github.com/exxxception/restful-service/models"
	"github.com/exxxception/restful-service/pkg/repository"
)

type AccountsService struct {
	repo repository.Accounts
}

func NewAccountsService(repo repository.Accounts) *AccountsService {
	return &AccountsService{
		repo: repo,
	}
}

func (s *AccountsService) Create(amount float64) (int, error) {
	return s.repo.Create(amount)
}

func (s *AccountsService) GetInfo(accountId int) (models.Account, error) {
	return s.repo.GetInfo(accountId)
}

func (s *AccountsService) Deposit(accountId int, amount float64) error {
	return s.repo.Deposit(accountId, amount)
}

func (s *AccountsService) Withdraw(accountId int, amount float64) error {
	return s.repo.Withdraw(accountId, amount)
}
