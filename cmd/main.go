package main

import (
	"bank/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/accounts", handler.CreateAccountHandler)
	r.POST("/accounts/:id/deposit", handler.DepositHandler)
	r.POST("/accounts/:id/withdraw", handler.WithdrawHandler)
	r.GET("/accounts/:id/balance", handler.GetBalanceHandler)

	r.Run(":8080")
}
