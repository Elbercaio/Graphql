package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type User struct {
	db   *sql.DB
	ID   string
	Name string
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) Create(name string) (User, error) {
	id := uuid.New().String()
	_, err := u.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		return User{}, err
	}
	return User{ID: id, Name: name}, nil
}