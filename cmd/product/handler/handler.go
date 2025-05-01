package handler

import (
	"fmt"
	"net/http"
	"produkfc/cmd/product/usecase"
	"produkfc/infrastructure/logger"
	"produkfc/models"
	"strconv"

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

func (h *ProductHandler) GetProductInfo(c *gin.Context) {
	param := c.Param("id")
	productID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"product id ": productID,
		}).Errorf("strconv.parse.int got error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "invalid product id",
		})
		return
	}
	product, err := h.ProductUseCase.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"product id ": productID,
		}).Errorf("h.ProductUseCase.GetProductBy ID got errror %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error message": err,
		})
		return
	}
	if product.ID == 0 {
		logger.Logger.WithFields(logrus.Fields{
			"product category id": productID,
		}).Info("Product not found")
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "productnot exists",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})

}
func (h *ProductHandler) GetProductCategoryInfo(c *gin.Context) {
	param := c.Param("id")
	productCategoryID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"product category id ": productCategoryID,
		}).Errorf("strconv.parse.int got error %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "invalid product category id",
		})
		return
	}
	productCategory, err := h.ProductUseCase.GetProductCategoryByID(c.Request.Context(), productCategoryID)
	if err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"product id ": productCategoryID,
		}).Errorf("h.ProductUseCase.GetProductCategoryByID got errror %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error message": err,
		})
		return
	}
	if productCategory.ID == 0 {
		logger.Logger.WithFields(logrus.Fields{
			"product category id": productCategoryID,
		}).Info("Product category not found")
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "product category not exists",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": productCategory,
	})

}

func (h *ProductHandler) ProductManagement(c *gin.Context) {
	var param = models.ProductParam{}
	if err := c.ShouldBindJSON(&param); err != nil {
		logger.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "invalid request",
		})
		return
	}

	if param.Action == "" {
		logger.Logger.Error("missing parameter action")
		c.JSON(http.StatusBadRequest, gin.H{
			"error message": "missing required parameter",
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
		productID, err := h.ProductUseCase.CreateNewProduct(c.Request.Context(), &param.Product)
		if err != nil {
			logger.Logger.WithFields(
				logrus.Fields{
					"param": param,
				}).Errorf("h.Product.CreateNewProduct got error %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error message": "err",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Succesfully create new product %d", productID),
		})

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
		product, err := h.ProductUseCase.EditProduct(c.Request.Context(), &param.Product)
		if err != nil {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUseCase.EditProduct got error %v ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully updated product",
			"product": product,
		})

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
		err := h.ProductUseCase.DeleteProduct(c.Request.Context(), param.ID)
		if err != nil {
			logger.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUseCaseDeleteProduct gor error %v ", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Product %d successfully deleted", param.ID),
		})

	default:
		logger.Logger.Errorf("Invalid action: %s", param.Action)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Action",
		})

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
