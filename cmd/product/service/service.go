package service

import (
	"context"
	"produkfc/cmd/product/repository"
	"produkfc/infrastructure/logger"
	"produkfc/models"

	"github.com/sirupsen/logrus"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(p repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: p,
	}
}

func (s *ProductService) GetProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	//get from redis
	product, err := s.ProductRepository.GetProductByIDFromRedis(ctx, productID)
	if err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"productID": productID,
		}).Errorf("s.ProductRepository.GetProductByIDFromRedis got error %v", err)
	}

	if product != nil && product.ID != 0 {
		return product, nil
	}
	//get from db
	product, err = s.ProductRepository.FindProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	//concurrent go routine
	type contextKey string
	const requestIDKey = contextKey("request_id")
	//ctxConcurrent := context.WithValue(ctx, context.Background(), ctx.Value("request_id"))
	ctxConcurrent := context.WithValue(ctx, requestIDKey, ctx.Value(requestIDKey))
	go func(ctx context.Context, product *models.Product, productID int64) {
		errConcurrent := s.ProductRepository.SetProductByID(ctx, product, productID)
		if errConcurrent != nil {
			logger.Logger.WithFields(logrus.Fields{
				"product": product,
			}).Errorf("s.ProductRepository.SetProductByID got error %v", errConcurrent)
		}
	}(ctxConcurrent, product, productID)
	return product, nil
}
func (s *ProductService) GetProductCategoryByID(ctx context.Context, productCategoryID int64) (*models.ProductCategory, error) {
	productCategory, err := s.ProductRepository.FindProductCategoryID(ctx, productCategoryID)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}

func (s *ProductService) CreateNewProduct(ctx context.Context, param *models.Product) (int64, error) {
	productID, err := s.ProductRepository.InsertNewProduct(ctx, param)
	if err != nil {
		return 0, err
	}
	return productID, nil
}
func (s *ProductService) CreateNewProductCategory(ctx context.Context, param *models.ProductCategory) (int64, error) {
	productCategoryID, err := s.ProductRepository.InsertNewCategoryProduct(ctx, param)
	if err != nil {
		return 0, err
	}
	return int64(productCategoryID), nil
}
func (s *ProductService) EditProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product, err := s.ProductRepository.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s *ProductService) EditProductCategory(ctx context.Context, productCategory *models.ProductCategory) (*models.ProductCategory, error) {
	productCategory, err := s.ProductRepository.UpdateProductCategory(ctx, productCategory)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}
func (s *ProductService) DeleteProduct(ctx context.Context, productID int64) error {
	err := s.ProductRepository.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProductService) DeleteProductCategory(ctx context.Context, productCategoryID int64) error {
	err := s.ProductRepository.DeleteProductCategory(ctx, productCategoryID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) SearchProduct(ctx context.Context, param models.SearchProductParameter) ([]models.Product, int, error) {
	product, totalCount, err := s.ProductRepository.SearchProduct(ctx, param)
	if err != nil {
		return nil, 0, err
	}
	return product, totalCount, nil
}
