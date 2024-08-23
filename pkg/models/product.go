package models

type Product struct {
    ProductID string `json:"product_id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Quantity float64  `json:"quantity"`
    Description string `json:"description"`
}
