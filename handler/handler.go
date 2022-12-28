package handler

import (
	"github.com/LastBit97/ewallet-restapi/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	walletService *service.WalletService
}

func NewHandler(walletService *service.WalletService) *Handler {
	return &Handler{walletService}
}

func (h *Handler) InitRoutes(rg *gin.RouterGroup) {
	rg.POST("/send")
	rg.GET("/wallet/{address}/balance")
}
