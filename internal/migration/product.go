package migration

import "database/sql"

func InitProductTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			description VARCHAR(255),
			tags TEXT[],
			quantity INT
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return InsertProducts(db)
}

func InsertProducts(db *sql.DB) error {
	query := `
		INSERT INTO products (description, tags, quantity)
		SELECT 'Product 1', '{tag1, tag2}', 10
		WHERE NOT EXISTS (SELECT 1 FROM products WHERE description = 'Product 1');

		INSERT INTO products (description, tags, quantity)
		SELECT 'Product 2', '{tag3, tag4}', 20
		WHERE NOT EXISTS (SELECT 1 FROM products WHERE description = 'Product 2');

		INSERT INTO products (description, tags, quantity)
		SELECT 'Product 3', '{tag5, tag6}', 30
		WHERE NOT EXISTS (SELECT 1 FROM products WHERE description = 'Product 3');
	`
	_, err := db.Exec(query)
	return err
}
