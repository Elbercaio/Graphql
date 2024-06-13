package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Todo struct {
	db     *sql.DB
	ID     string
	Text   string
	Done   bool
	UserID string
}

func NewTodo(db *sql.DB) *Todo {
	return &Todo{db: db}
}

func (t *Todo) Create(text string, userId string) (*Todo, error) {
	id := uuid.New().String()
	_, err := t.db.Exec("INSERT INTO todos (id, text, user_id) VALUES ($1, $2, $3)",
	id, text, userId)
	if err != nil {
		return &Todo{}, err
	}
	return &Todo{ID: id, Text: text, UserID: userId}, nil
}

func (t *Todo) List() ([]Todo, error) {
	rows, err := t.db.Query("SELECT id, text, done, user_id FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var id string
		var text string
		var done bool
		var userId string
		if err := rows.Scan(&id, &text, &done, &userId); err != nil {
			return nil, err
		}
		todos = append(todos, Todo{ID: id, Text: text, Done: done, UserID: userId})
	}
	return todos, nil
}

func (t *Todo) Get(id string) (Todo, error) {
	row := t.db.QueryRow("SELECT text, done FROM todos WHERE id = $1", id)
	var text string
	var done bool
	if err := row.Scan(&text, &done); err != nil {
		return Todo{}, err
	}
	return Todo{ID: id, Text: text, Done: done}, nil
}

func (t *Todo) GetByUserId(userId string) ([]Todo, error) {
	rows, err := t.db.Query("SELECT id, text, done FROM todos WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var id string
		var text string
		var done bool
		if err := rows.Scan(&id, &text, &done); err != nil {
			return nil, err
		}
		todos = append(todos, Todo{ID: id, Text: text, Done: done})
	}
	return todos, nil
}