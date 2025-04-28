package usecase

import (
	"context"
	"produkfc/cmd/product/service"
	"produkfc/infrastructure/logger"
	"produkfc/models"

	"github.com/sirupsen/logrus"
)

type ProductUsecase struct {
	ProductService service.ProductService
}

func NewProductUsecase(ps service.ProductService) *ProductUsecase {
	return &ProductUsecase{
		ProductService: ps,
	}
}

func (uc *ProductUsecase) GetProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	product, err := uc.ProductService.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (uc *ProductUsecase) GetProductCategoryByID(ctx context.Context, productCategoryID int64) (*models.ProductCategory, error) {
	productCategory, err := uc.ProductService.GetProductCategoryByID(ctx, productCategoryID)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}

func (uc *ProductUsecase) CreateNewProduct(ctx context.Context, param *models.Product) (int64, error) {
	productID, err := uc.ProductService.CreateNewProduct(ctx, param)
	if err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"name":     param.Name,
			"category": param.CategoryID,
		}).Errorf("uc.ProductService.CreateNewProduct got error %v", err)
		return 0, nil
	}

	return productID, nil
}
func (uc *ProductUsecase) CreateNewProductCategory(ctx context.Context, param *models.ProductCategory) (int64, error) {
	productCategoryID, err := uc.ProductService.CreateNewProductCategory(ctx, param)
	if err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"name": param.Name,
		}).Errorf("uc.ProductService.CreateNewProductCategory got error %v ", err)
		return 0, nil
	}

	return productCategoryID, nil
}
func (uc *ProductUsecase) EditProduct(ctx context.Context, param *models.Product) (*models.Product, error) {
	product, err := uc.ProductService.EditProduct(ctx, param)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (uc *ProductUsecase) EditProductCategory(ctx context.Context, param *models.ProductCategory) (*models.ProductCategory, error) {
	productCategory, err := uc.ProductService.EditProductCategory(ctx, param)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}
func (uc *ProductUsecase) DeleteProduct(ctx context.Context, productID int64) error {
	err := uc.ProductService.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}
	return nil
}
func (uc *ProductUsecase) DeleteProductCategory(ctx context.Context, productCategoryID int64) error {
	err := uc.ProductService.DeleteProductCategory(ctx, productCategoryID)
	if err != nil {
		return err
	}
	return nil
}
