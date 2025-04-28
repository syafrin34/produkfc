package handler

import (
	"net/http"
	"produkfc/cmd/product/usecase"
	"produkfc/infrastructure/logger"
	"produkfc/models"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUseCase usecase.ProductUsecase
}

func NewProductHandler(pu usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUseCase: pu,
	}
}
func (h *ProductHandler) ProductCategoryManagement(c *gin.Context) {
	var param models.ProductCategoryParam
	if err := c.ShouldBindJSON(&param); err != nil {
		logger.Logger.Error(err.Error()) // untuk debugging
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Input",
		})
		return
	}

	if param.Action == "" {
		logger.Logger.Error("missing parameter action") // untuk debugging
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Missing required paramete",
		})
		return
	}

	switch param.Action {
	case "add":
	case "edit":
	case "delete":
	default:
		logger.Logger.Errorf("Invalid action: %s", param.Action)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Action",
		})
		return
	}
}
