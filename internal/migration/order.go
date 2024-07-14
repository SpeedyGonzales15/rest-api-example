package migration

import "database/sql"

func InitOrderTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT,
		products JSONB
	);`
	_, err := db.Exec(query)
	return err
}
