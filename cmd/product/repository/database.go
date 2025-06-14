package repository

import (
	"context"
	"errors"
	"fmt"
	"produkfc/models"

	"gorm.io/gorm"
)

func (r *ProductRepository) FindProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	var product models.Product
	err := r.Database.WithContext(ctx).Table("product").Where("id = ?", productID).Last(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.Product{}, nil
		}

		return nil, err
	}
	return &product, nil

}

func (r *ProductRepository) FindProductCategoryID(ctx context.Context, productCategoryID int64) (*models.ProductCategory, error) {
	var productCategory models.ProductCategory
	err := r.Database.WithContext(ctx).Table("product_category").Where("id = ?", productCategoryID).Last(&productCategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ProductCategory{}, nil
		}
		return nil, err
	}
	return &productCategory, nil
}

func (r *ProductRepository) InsertNewProduct(ctx context.Context, product *models.Product) (int64, error) {
	err := r.Database.WithContext(ctx).Table("product").Create(product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (r *ProductRepository) InsertNewCategoryProduct(ctx context.Context, productCategory *models.ProductCategory) (int, error) {
	err := r.Database.WithContext(ctx).Table("product_category").Create(productCategory).Error
	if err != nil {
		return 0, err
	}
	return productCategory.ID, nil
}
func (r *ProductRepository) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	err := r.Database.WithContext(ctx).Table("product").Save(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProductCategory(ctx context.Context, productCategory *models.ProductCategory) (*models.ProductCategory, error) {
	err := r.Database.WithContext(ctx).Table("product_category").Save(productCategory).Error
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, productID int64) error {
	err := r.Database.WithContext(ctx).Table("product").Delete(&models.Product{}, productID).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *ProductRepository) DeleteProductCategory(ctx context.Context, productCategoryID int64) error {
	err := r.Database.WithContext(ctx).Table("product_category").Delete(&models.ProductCategory{}, productCategoryID).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *ProductRepository) SearchProduct(ctx context.Context, param models.SearchProductParameter) ([]models.Product, int, error) {
	var products []models.Product
	var totalCount int64
	query := r.Database.WithContext(ctx).Table("product").
		Select("product.id, product.name, product.description, product.price, product.stock,product.category_id, product_category.name AS category").
		Joins("JOIN product_category ON product.category_id=product_category.id")
	// filtering
	if param.Name != "" {
		query = query.Where("product.name ILIKE ?", "%"+param.Name+"%")
	}
	if param.Category != "" {
		query = query.Where("product_category.name = ?", param.Category)
	}
	if param.MinPrice > 0 {
		query = query.Where("product.price >= ?", param.MinPrice)
	}
	if param.MaxPrice > 0 {
		query = query.Where("product.price <= ?", param.MaxPrice)
	}

	//pagination

	//dapetin total counts dari query
	query.Model(&models.Product{}).Count(&totalCount)

	//default orderBy
	if param.OrderBy == "" {
		param.OrderBy = "product.name"
	}
	if param.Sort == "" || (param.Sort != "ASC" && param.Sort != "DESC") {
		param.Sort = "ASC"
	}
	orderBy := fmt.Sprintf("%s %s", param.OrderBy, param.Sort)
	query = query.Order(orderBy)

	offset := (param.Page - 1) * param.PageSize
	query = query.Limit(param.PageSize).Offset(offset)

	err := query.Scan(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, int(totalCount), err

}

func (r *ProductRepository) DeductProductStockByProductID(ctx context.Context, productID int64, qty int) error {
	err := r.Database.Table("product").WithContext(ctx).Model(&models.Product{}).
		Updates(map[string]interface{}{
			"stock": gorm.Expr("stock - %d", qty),
		}).Where("id = ?", productID).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) AddProductStockByProductID(ctx context.Context, productID int64, qty int) error {
	err := r.Database.Table("products").WithContext(ctx).Model(&models.Product{}).
		Updates(map[string]interface{}{
			"stock": gorm.Expr("stock + %d", qty),
		}).Where("id = ? ", productID).Error
	if err != nil {
		return err
	}
	return nil
}
