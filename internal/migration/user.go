package migration

import (
	"database/sql"
)

func InitUserTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			full_name VARCHAR(255),
			age INT,
			is_married BOOLEAN,
			password VARCHAR(255)
		);
	`
	_, err := db.Exec(query)
	return err
}
