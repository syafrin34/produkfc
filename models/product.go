package models

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  int     `json:"category_id"`
}

type ProductCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductCategoryParam struct {
	Action string `json:"action"`
	ProductCategory
}

type ProductParam struct {
	Action string `json:"action"`
	Product
}

type SearchProductParameter struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	MinPrice float64 `json:"minPrice"`
	MaxPrice float64 `json:"maxPrice"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
	OrderBy  string  `json:"orderBy"`
	Sort     string  `json:"sort"`
}
