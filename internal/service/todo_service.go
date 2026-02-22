package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"todo-app/internal/model"
	"todo-app/internal/repository"

	"github.com/google/uuid"
)

var ErrInvalidTask = errors.New("task cannot be empty")

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(r repository.TodoRepository) *TodoService {
	return &TodoService{repo: r}
}

func (s *TodoService) Create(ctx context.Context, task string, due time.Time) (model.Todo, error) {
	if strings.TrimSpace(task) == "" {
		return model.Todo{}, ErrInvalidTask
	}

	todo := model.Todo{
		ID:        uuid.New().String(),
		Task:      task,
		DueDate:   due,
		Completed: false,
		CreatedAt: time.Now().UTC(),
	}

	return todo, s.repo.Create(ctx, todo)
}

func (s *TodoService) Get(ctx context.Context, id string) (model.Todo, error) {
	return s.repo.Get(ctx, id)
}

func (s *TodoService) Update(ctx context.Context, t model.Todo) error {
	if strings.TrimSpace(t.Task) == "" {
		return ErrInvalidTask
	}
	return s.repo.Update(ctx, t)
}

func (s *TodoService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *TodoService) List(ctx context.Context, include bool) ([]model.Todo, error) {
	return s.repo.List(ctx, include)
}
