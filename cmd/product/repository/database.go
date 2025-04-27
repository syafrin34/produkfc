package repository

import (
	"context"
	"produkfc/models"
)

func (r *ProductRepository) FindProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	var product models.Product
	err := r.Database.WithContext(ctx).Table("product").Where("id = ?", productID).Last(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (r *ProductRepository) FindProductCategoryID(ctx context.Context, productCategoryID int64) (*models.ProductCategory, error) {
	var productCategory models.ProductCategory
	err := r.Database.WithContext(ctx).Table("category_product").Where("id = ?", productCategoryID).Last(&productCategory).Error
	if err != nil {
		return nil, err
	}
	return &productCategory, nil
}
