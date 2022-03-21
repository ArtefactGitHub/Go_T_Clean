package interactor

import (
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
	"github.com/ArtefactGitHub/Go_T_Clean/infurastructure"
)

type TaskInteractor struct {
	repository *infurastructure.InMemoryTaskRepository
}

func NewTaskInteractor() TaskInteractor {
	r := infurastructure.NewTaskRepository()
	return TaskInteractor{repository: &r}
}

func (i TaskInteractor) GetAll() ([]model.Task, error) {
	return i.repository.GetAll()
}

func (i TaskInteractor) Get(id int) (*model.Task, error) {
	return i.repository.Get(id)
}

func (i TaskInteractor) Create(task model.Task) (int, error) {
	return i.repository.Create(task)
}

func (i TaskInteractor) Update(task model.Task) (*model.Task, error) {
	return i.repository.Update(task)
}

func (i TaskInteractor) Delete(id int) (bool, error) {
	return i.repository.Delete(id)
}
