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

func (u *User) Create(name string) (*User, error) {
	id := uuid.New().String()
	_, err := u.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		return &User{}, err
	}
	return &User{ID: id, Name: name}, nil
}

func (u *User) List() ([]User, error) {
	rows, err := u.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var id string
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		users = append(users, User{ID: id, Name: name})
	}
	return users, nil
}

func (u *User) Get(id string) (User, error) {
	row := u.db.QueryRow("SELECT name FROM users WHERE id = $1", id)
	var name string
	if err := row.Scan(&name); err != nil {
		return User{}, err
	}
	return User{ID: id, Name: name}, nil
}

func (u *User) GetByTodoId(todoId string) (User, error) {
	var id string
	var name string
	err := u.db.QueryRow("SELECT id, name FROM users INNER JOIN todos on todos.user_id = users.id WHERE todos.id = $1", todoId).Scan(&id, &name)
	if err != nil {
		return User{}, err
	}
	return User{ID: id, Name: name}, nil
}