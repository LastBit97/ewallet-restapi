package handler

import (
	"net/http"

	"github.com/LastBit97/ewallet-restapi/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	walletService service.WalletService
}

func NewHandler(walletService service.WalletService) Handler {
	return Handler{walletService}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/send")
	rg.GET("/wallet/:address/balance", h.GetBalance)
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
