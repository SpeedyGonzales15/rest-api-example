package models

type Product struct {
	ID          int      `json:"id" db:"id"`
	Description string   `json:"description" db:"description"`
	Tags        []string `json:"tags" db:"tags"`
	Quantity    int      `json:"quantity" db:"quantity"`
}
