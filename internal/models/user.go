package models

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	FullName  string `json:"full_name" db:"full_name"`
	Age       int    `json:"age" db:"age"`
	IsMarried bool   `json:"is_married" db:"is_married"`
	Password  string `json:"password" db:"password"`
}

type UpdateUser struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	FullName  *string `json:"full_name"`
	Age       *int    `json:"age"`
	IsMarried *bool   `json:"is_married"`
	Password  *string `json:"password"`
}
