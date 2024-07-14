package repository

import (
	"database/sql"
	"fmt"
	"rest-api-example/internal/models"
	"rest-api-example/pkg/middleware"
	"strings"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user models.User) (int, error) {
	var id int
	query := "INSERT INTO users (first_name, last_name, full_name, age, is_married, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.db.QueryRow(query, user.FirstName, user.LastName, fmt.Sprint(user.FirstName+" "+user.LastName), user.Age, user.IsMarried, middleware.PasswordHash(user.Password))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserPostgres) GetAll() ([]models.User, error) {
	var list []models.User

	query := "SELECT id, first_name, last_name, full_name, age, is_married FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.FullName, &user.Age, &user.IsMarried)
		if err != nil {
			return nil, err
		}
		list = append(list, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *UserPostgres) GetById(id int) (models.User, error) {
	var user models.User
	query := "SELECT id, first_name, last_name, full_name, age, is_married FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.FullName, &user.Age, &user.IsMarried); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserPostgres) Update(userId int, user models.UpdateUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argId))
		args = append(args, *user.FirstName)
		argId++
	}

	if user.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, *user.LastName)
		argId++
	}

	if user.FirstName != nil || user.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("full_name=$%d", argId))
		args = append(args, *user.FirstName+" "+*user.LastName)
		argId++
	}

	if user.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *user.Age)
		argId++
	}

	if user.IsMarried != nil {
		setValues = append(setValues, fmt.Sprintf("is_married=$%d", argId))
		args = append(args, *user.IsMarried)
		argId++
	}

	if user.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *user.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *UserPostgres) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
