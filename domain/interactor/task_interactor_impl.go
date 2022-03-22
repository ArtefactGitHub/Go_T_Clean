package interactor

import (
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
	"github.com/ArtefactGitHub/Go_T_Clean/infurastructure"
	"github.com/ArtefactGitHub/Go_T_Clean/usecase/interfaces"
)

type taskInteractor struct {
	repository *infurastructure.InMemoryTaskRepository
}

func NewTaskInteractor() interfaces.TaskInteractor {
	r := infurastructure.NewTaskRepository()
	return taskInteractor{repository: &r}
}

func (i taskInteractor) GetAll() ([]model.Task, error) {
	return i.repository.GetAll()
}

func (i taskInteractor) Get(id int) (*model.Task, error) {
	return i.repository.Get(id)
}

func (i taskInteractor) Create(task model.Task) (int, error) {
	return i.repository.Create(task)
}

func (i taskInteractor) Update(task model.Task) (*model.Task, error) {
	return i.repository.Update(task)
}

func (i taskInteractor) Delete(id int) (bool, error) {
	return i.repository.Delete(id)
}
