package repository

import (
	"database/sql"
	"rest-api-example/internal/models"

	"github.com/lib/pq"
)

type ProductPostgres struct {
	db *sql.DB
}

func NewProductPostgres(db *sql.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) GetAllProducts() ([]models.Product, error) {
	var list []models.Product

	query := "SELECT id, description, tags, quantity FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Description, pq.Array(&product.Tags), &product.Quantity)
		if err != nil {
			return nil, err
		}
		list = append(list, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *ProductPostgres) GetProductsById(id int) (models.Product, error) {
	var product models.Product

	query := "SELECT id, description, tags, quantity FROM products WHERE id = $1"
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&product.ID, &product.Description, pq.Array(&product.Tags), &product.Quantity); err != nil {
		return models.Product{}, err
	}
	return product, nil
}
