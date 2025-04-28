package handler

import (
	"fmt"
	"net/http"
	"produkfc/cmd/product/usecase"
	"produkfc/infrastructure/logger"
	"produkfc/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		if param.ID != 0 {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("invalid request - product category id is not empty")
			c.JSON(http.StatusBadRequest, gin.H{
				"error message": "Invalid Request",
			})
			return
		}
		productCategoryID, err := h.ProductUseCase.CreateNewProductCategory(c.Request.Context(), &param.ProductCategory)
		if err != nil {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("h.ProductUseCase.CreateNewProductCategory got error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Succesfully create new product category: %d", &productCategoryID),
		})
		return
	case "edit":
		if param.ID == 0 {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("invalid request - product id is empty")
			c.JSON(http.StatusBadRequest, gin.H{
				"error message": "Invalid Request",
			})
			return
		}
		productCategory, err := h.ProductUseCase.EditProductCategory(c.Request.Context(), &param.ProductCategory)
		if err != nil {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUseCase.EdiProductCategory got error %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":         "Succesfully edit product",
			"productCategory": productCategory,
		})
		return
	case "delete":
		if param.ID == 0 {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("invalid request - product id is empty")
			c.JSON(http.StatusBadRequest, gin.H{
				"error message": "Invalid Request",
			})
			return
		}
		err := h.ProductUseCase.DeleteProductCategory(c.Request.Context(), int64(param.ID))
		if err != nil {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUseCase.DeleteProductCategory got error %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Product Category ID %d succesfully deleted", param.ID),
		})
		return
	default:
		logger.Logger.Errorf("Invalid action: %s", param.Action)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Action",
		})
		return
	}
}
