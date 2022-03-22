package infurastructure

import (
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
	"github.com/ArtefactGitHub/Go_T_Clean/usecase/interfaces"
)

type InMemoryTaskRepository struct {
	tasks []model.Task
}

func NewInMemoryTaskRepository() interfaces.TaskRepository {
	// 仮データ
	tasks := []model.Task{model.NewTask(0, "first")}
	r := InMemoryTaskRepository{
		tasks: tasks,
	}
	return &r
}

func (r *InMemoryTaskRepository) GetAll() ([]model.Task, error) {
	return r.tasks, nil
}

func (r *InMemoryTaskRepository) Get(id int) (*model.Task, error) {
	exist := r.get(id)
	if exist != nil {
		result := model.NewTask(exist.Id, exist.Name)
		return &result, nil
	}

	return nil, nil
}

func (r *InMemoryTaskRepository) Create(task model.Task) (int, error) {
	id := r.createNewId()
	newTask := model.NewTask(id, task.Name)
	r.tasks = append(r.tasks, newTask)
	return id, nil
}

func (r *InMemoryTaskRepository) Update(task model.Task) (*model.Task, error) {
	for _, v := range r.tasks {
		if v.Id == task.Id {
			r.tasks[v.Id].Name = task.Name
			result := model.NewTask(task.Id, task.Name)
			return &result, nil
		}
	}

	return nil, nil
}

func (r *InMemoryTaskRepository) Delete(id int) (bool, error) {
	index := r.getIndex(id)
	if index < 0 {
		return false, nil
	}

	r.tasks = r.remove(index)
	return true, nil
}

func (r *InMemoryTaskRepository) createNewId() int {
	if len(r.tasks) > 0 {
		return r.tasks[len(r.tasks)-1].Id + 1
	} else {
		return 0
	}
}

func (r *InMemoryTaskRepository) get(id int) *model.Task {
	for _, v := range r.tasks {
		if v.Id == id {
			return &v
		}
	}

	return nil
}

func (r *InMemoryTaskRepository) getIndex(id int) int {
	for i, v := range r.tasks {
		if v.Id == id {
			return i
		}
	}

	return -1
}

func (r *InMemoryTaskRepository) remove(index int) []model.Task {
	return append(r.tasks[:index], r.tasks[index+1:]...)
}
