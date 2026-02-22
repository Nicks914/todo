package repository

import (
	"context"
	"database/sql"
	"errors"

	"todo-app/internal/model"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Create(ctx context.Context, todo model.Todo) error {
	query := `
	INSERT INTO todos (id, task, due_date, completed, created_at)
	VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query,
		todo.ID, todo.Task, todo.DueDate, todo.Completed, todo.CreatedAt)
	return err
}

func (r *MySQLRepository) Get(ctx context.Context, id string) (model.Todo, error) {
	var t model.Todo
	query := `SELECT id, task, due_date, completed, created_at FROM todos WHERE id = ?`

	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&t.ID, &t.Task, &t.DueDate, &t.Completed, &t.CreatedAt)

	if err == sql.ErrNoRows {
		return model.Todo{}, errors.New("todo not found")
	}
	return t, err
}

func (r *MySQLRepository) Update(ctx context.Context, todo model.Todo) error {
	query := `
	UPDATE todos SET task=?, due_date=?, completed=?
	WHERE id=?
	`
	res, err := r.db.ExecContext(ctx, query,
		todo.Task, todo.DueDate, todo.Completed, todo.ID)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (r *MySQLRepository) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx,
		"DELETE FROM todos WHERE id=?", id)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (r *MySQLRepository) List(ctx context.Context, includeCompleted bool) ([]model.Todo, error) {
	query := `
	SELECT id, task, due_date, completed, created_at
	FROM todos
	WHERE (? = TRUE OR completed = FALSE)
	ORDER BY due_date ASC
	`

	rows, err := r.db.QueryContext(ctx, query, includeCompleted)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo

	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Task, &t.DueDate, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}
