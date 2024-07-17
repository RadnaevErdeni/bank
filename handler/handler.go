package handler

import (
	bank "bank"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAccountHandler(c *gin.Context) {
	account := bank.CreateAccount()
	c.JSON(http.StatusOK, gin.H{"id": account.ID})
}

func DepositHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	account, exists := bank.GetAccount(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	if err := account.Deposit(request.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deposit successful"})
}

func WithdrawHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	account, exists := bank.GetAccount(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	if err := account.Withdraw(request.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "withdraw successful"})
}

func GetBalanceHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	account, exists := bank.GetAccount(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	balance := account.GetBalance()
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
