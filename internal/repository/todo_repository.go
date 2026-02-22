package repository

import (
	"context"
	"todo-app/internal/model"
)

type TodoRepository interface {
	Create(ctx context.Context, todo model.Todo) error
	Get(ctx context.Context, id string) (model.Todo, error)
	Update(ctx context.Context, todo model.Todo) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, includeCompleted bool) ([]model.Todo, error)
}
