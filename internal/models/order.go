package models

type ProductItem struct {
	ProductId       int      `json:"product_id" db:"product_id"`
	Description     string   `json:"description"`
	Tags            []string `json:"tags"`
	QuantityInOrder int      `json:"quantity_in_order"`
}
type Order struct {
	ID       int           `json:"id" db:"id"`
	UserId   int           `json:"user_id" db:"user_id"`
	Products []ProductItem `json:"products" db:"products"`
}
