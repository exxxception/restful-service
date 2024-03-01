package postgres

import (
	"fmt"

	"github.com/exxxception/restful-service/models"
)

type AccountsPostgres struct {
	postgres *PostgresDB
}

func NewAccountsPostgres(pos *PostgresDB) *AccountsPostgres {
	return &AccountsPostgres{postgres: pos}
}

func (r *AccountsPostgres) Create(amount float64) (int, error) {
	var accountId int

	createAccountQuery := fmt.Sprintf(`
	INSERT INTO %s(balance) values ($1) RETURNING id
	`, accountsTable)
	if err := r.postgres.db.QueryRow(createAccountQuery, amount).Scan(&accountId); err != nil {
		return 0, err
	}

	return accountId, nil
}

func (r *AccountsPostgres) GetInfo(accountId int) (models.Account, error) {
	var account models.Account

	getAccountInfoQuery := fmt.Sprintf(`
	SELECT id, balance 
	FROM %s
	WHERE id = $1
	`, accountsTable)
	err := r.postgres.db.QueryRow(getAccountInfoQuery, accountId).Scan(&account.Id, &account.Balance)

	return account, err
}

func (r *AccountsPostgres) Deposit(accountId int, amount float64) error {
	depositQuery := fmt.Sprintf(`
	UPDATE %s SET balance = balance + $1 WHERE id = $2
	`, accountsTable)
	_, err := r.postgres.db.Exec(depositQuery, amount, accountId)
	return err
}

func (r *AccountsPostgres) Withdraw(accountId int, amount float64) error {
	withdrawQuery := fmt.Sprintf(`
	UPDATE %s SET balance = balance - $1 WHERE id = $2
	`, accountsTable)
	_, err := r.postgres.db.Exec(withdrawQuery, amount, accountId)
	return err
}
