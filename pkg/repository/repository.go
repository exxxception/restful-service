package repository

import (
	"log"
	"os"

	"github.com/exxxception/restful-service/models"
	"github.com/exxxception/restful-service/pkg/repository/postgres"
)

type Accounts interface {
	Create(amount float64) (int, error)
	GetInfo(accountId int) (models.Account, error)
	Deposit(accountId int, amount float64) error
	Withdraw(accountId int, amount float64) error
}

type Repository struct {
	Accounts
}

func NewRepository() *Repository {
	repo, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return &Repository{
		Accounts: postgres.NewAccountsPostgres(repo),
	}
}
