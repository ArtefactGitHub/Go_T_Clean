package interactor

import (
	"context"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
	"github.com/ArtefactGitHub/Go_T_Clean/usecase/interfaces"
)

type taskInteractor struct {
	repository interfaces.TaskRepository
}

func NewTaskInteractor(repository interfaces.TaskRepository) interfaces.TaskInteractor {
	return taskInteractor{repository: repository}
}

func (i taskInteractor) GetAll(ctx context.Context) ([]model.Task, error) {
	return i.repository.GetAll(ctx)
}

func (i taskInteractor) Get(ctx context.Context, id int) (*model.Task, error) {
	return i.repository.Get(ctx, id)
}

func (i taskInteractor) Create(ctx context.Context, task model.Task) (int, error) {
	return i.repository.Create(ctx, task)
}

func (i taskInteractor) Update(ctx context.Context, task model.Task) (*model.Task, error) {
	return i.repository.Update(ctx, task)
}

func (i taskInteractor) Delete(ctx context.Context, id int) (bool, error) {
	return i.repository.Delete(ctx, id)
}
