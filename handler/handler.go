package handler

import (
	"net/http"
	"strconv"

	"github.com/LastBit97/ewallet-restapi/model"
	"github.com/LastBit97/ewallet-restapi/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	walletService      service.WalletService
	transactionService service.TransactionService
}

func NewHandler(walletService service.WalletService, transactionService service.TransactionService) Handler {
	return Handler{walletService, transactionService}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/send", h.Send)
	rg.GET("/wallet/:address/balance", h.GetBalance)
	rg.GET("/transactions", h.GetLast)
}

func (h *Handler) GetBalance(ctx *gin.Context) {
	walletAddress := ctx.Param("address")
	wallet, err := h.walletService.GetWallet(walletAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": wallet.Balance})
}

func (h *Handler) Send(ctx *gin.Context) {
	var transaction *model.CreateTransactionRequest

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newTransaction, err := h.transactionService.Send(transaction)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newTransaction})
}

func (h *Handler) GetLast(ctx *gin.Context) {
	count := ctx.DefaultQuery("count", "10")
	intCount, err := strconv.Atoi(count)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	transactions, err := h.transactionService.GetTransactions(intCount)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(transactions), "data": transactions})
}
