package task

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model/task"
)

type TaskInteractor interface {
	GetAll(ctx context.Context) ([]task.Task, error)
	Get(ctx context.Context, id int) (*task.Task, error)
	Create(ctx context.Context, task task.Task) (int, error)
	Update(ctx context.Context, task task.Task) (*task.Task, error)
	Delete(ctx context.Context, id int) (bool, error)
}

type taskInteractorImpl struct {
	repository task.TaskRepository
}

func NewTaskInteractor(repository task.TaskRepository) TaskInteractor {
	return taskInteractorImpl{repository: repository}
}

func (i taskInteractorImpl) GetAll(ctx context.Context) ([]task.Task, error) {
	return i.repository.GetAll(ctx)
}

func (i taskInteractorImpl) Get(ctx context.Context, id int) (*task.Task, error) {
	return i.repository.Get(ctx, id)
}

func (i taskInteractorImpl) Create(ctx context.Context, task task.Task) (int, error) {
	task.CreatedAt = time.Now()
	return i.repository.Create(ctx, task)
}

func (i taskInteractorImpl) Update(ctx context.Context, task task.Task) (*task.Task, error) {
	task.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return i.repository.Update(ctx, task)
}

func (i taskInteractorImpl) Delete(ctx context.Context, id int) (bool, error) {
	return i.repository.Delete(ctx, id)
}
