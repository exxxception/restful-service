package handlers

import (
	"net/http"
	"strconv"

	"github.com/exxxception/restful-service/models"
	"github.com/gin-gonic/gin"
)

// create creates a user account with provided amount of money
// @Summary Create a user account with given balance
// @Security		ApiKeyAuth
// @Tags    Accounts
// @ID account-create
// @Accept json
// @Param amount body models.CreateAccount true "amount of money on the new account"
// @Success 201 "Created" {json}
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router  /api/accounts [post]
func (h *Handlers) Create(c *gin.Context) {
	var amount models.CreateAccount

	if err := c.BindJSON(&amount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Accounts.Create(amount.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// getInfo gets a user account by id
// @Summary Get a user account information by its ID
// @Security		ApiKeyAuth
// @Tags    Accounts
// @ID account-get
// @Produce json
// @Param id path string true "account id"
// @Success 200 {object} models.Account "Account info"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router  /api/accounts/{id} [get]
func (h *Handlers) GetInfo(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	account, err := h.services.Accounts.GetInfo(accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

// depositFunds deposit funds to a given account
// @Summary Deposit funds to an account and add en entry to the transfers table
// @Security		ApiKeyAuth
// @Tags    Accounts
// @ID account-deposit
// @Accept json
// @Param id path string true "account id"
// @Param amount body models.DepositFunds true "amount of money"
// @Success 204 "No Content"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router /api/accounts/deposit/{id} [put]
func (h *Handlers) Deposit(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var amount models.DepositFunds
	if err := c.BindJSON(&amount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.Accounts.Deposit(accountId, amount.Amount); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// depositFunds withdraw funds from a given account
// @Summary Withdraw funds from an account and add en entry to the transfers table
// @Security		ApiKeyAuth
// @Tags    Accounts
// @ID account-withdraw
// @Accept json
// @Param id path string true "account id"
// @Param amount body models.WithdrawFunds true "amount of money"
// @Success 204 "No Content"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router /api/accounts/withdraw/{id} [put]
func (h *Handlers) Withdraw(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var amount models.WithdrawFunds
	if err := c.BindJSON(&amount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.Accounts.Withdraw(accountId, amount.Amount); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
