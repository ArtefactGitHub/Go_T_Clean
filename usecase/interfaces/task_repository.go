package interfaces

import (
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
)

type TaskRepository interface {
	GetAll() ([]model.Task, error)
	Get(id int) (*model.Task, error)
	Create(task model.Task) (int, error)
	Update(task model.Task) (*model.Task, error)
	Delete(id int) (bool, error)
}
