package task

import (
	"context"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]Task, error)
	Get(ctx context.Context, id int) (*Task, error)
	Create(ctx context.Context, task Task) (int, error)
	Update(ctx context.Context, task Task) (*Task, error)
	Delete(ctx context.Context, id int) (bool, error)
	Finalize()
}
