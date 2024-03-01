package models

type CreateAccount struct {
	Amount float64 `json:"amount"`
}

type DepositFunds struct {
	Amount float64 `json:"amount"`
}

type WithdrawFunds struct {
	Amount float64 `json:"amount"`
}
