package handler

import (
	"gotesting/Desafio/internal/product"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	sv product.Service
}

func NewHandler(sv product.Service) *Handler {
	return &Handler{
		sv: sv,
	}
}

func (h *Handler) GetProductsBySellerID(ctx *gin.Context) {
	sellerID := ctx.Query("seller_id")
	if strings.TrimSpace(sellerID) == "" {
		ctx.JSON(400, gin.H{"error": "seller_id query param is required"})
		return
	}

	products, err := h.sv.GetAllBySeller(sellerID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, products)
}
