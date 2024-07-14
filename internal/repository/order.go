package repository

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"rest-api-example/internal/models"

	"github.com/lib/pq"
)

type OrderPostgres struct {
	db *sql.DB
}

func NewOrderPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateOrder(order models.Order) (int, error) {
	var id int
	query := `
		INSERT INTO orders (user_id, products) VALUES ($1, $2::jsonb) RETURNING id;
	`

	products, err := json.Marshal(order.Products)
	if err != nil {
		return 0, err
	}

	row := r.db.QueryRow(query, order.UserId, string(products))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	for _, item := range order.Products {
		if err := r.decrementProductQuantity(item.ProductId, item.QuantityInOrder); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *OrderPostgres) decrementProductQuantity(productId int, quantity int) error {
	query := `
		UPDATE products SET quantity = quantity - $2 WHERE id = $1;
	`

	_, err := r.db.Exec(query, productId, quantity)

	return err
}

func (r *OrderPostgres) GetAllOrders(userId int) ([]models.Order, error) {
	var list []models.Order

	query := "SELECT id, user_id, products FROM orders WHERE user_id = $1"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserId, pq.Array(&order.Products)); err != nil {
			return nil, err
		}
		list = append(list, order)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}
