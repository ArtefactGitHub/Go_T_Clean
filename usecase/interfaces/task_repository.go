package interfaces

import (
	"context"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	Get(ctx context.Context, id int) (*model.Task, error)
	Create(ctx context.Context, task model.Task) (int, error)
	Update(ctx context.Context, task model.Task) (*model.Task, error)
	Delete(ctx context.Context, id int) (bool, error)
	Finalize()
}
