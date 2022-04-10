package infurastructure

import (
	"context"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model/task"
)

type inMemoryTaskRepository struct {
	tasks []task.Task
}

func NewInMemoryTaskRepository() (task.TaskRepository, error) {
	// 仮データ
	tasks := []task.Task{task.NewTask(0, "first")}
	r := inMemoryTaskRepository{
		tasks: tasks,
	}
	return &r, nil
}

func (r *inMemoryTaskRepository) Finalize() {}

func (r *inMemoryTaskRepository) GetAll(ctx context.Context) ([]task.Task, error) {
	return r.tasks, nil
}

func (r *inMemoryTaskRepository) Get(ctx context.Context, id int) (*task.Task, error) {
	exist := r.get(id)
	if exist != nil {
		result := task.NewTask(exist.Id, exist.Name)
		return &result, nil
	}

	return nil, nil
}

func (r *inMemoryTaskRepository) Create(ctx context.Context, t task.Task) (int, error) {
	id := r.createNewId()
	newTask := task.NewTask(id, t.Name)
	r.tasks = append(r.tasks, newTask)
	return id, nil
}

func (r *inMemoryTaskRepository) Update(ctx context.Context, t task.Task) (*task.Task, error) {
	for _, v := range r.tasks {
		if v.Id == t.Id {
			r.tasks[v.Id].Name = t.Name
			result := task.NewTask(t.Id, t.Name)
			return &result, nil
		}
	}

	return nil, nil
}

func (r *inMemoryTaskRepository) Delete(ctx context.Context, id int) (bool, error) {
	index := r.getIndex(id)
	if index < 0 {
		return false, nil
	}

	r.tasks = r.remove(index)
	return true, nil
}

func (r *inMemoryTaskRepository) createNewId() int {
	if len(r.tasks) > 0 {
		return r.tasks[len(r.tasks)-1].Id + 1
	} else {
		return 0
	}
}

func (r *inMemoryTaskRepository) get(id int) *task.Task {
	for _, v := range r.tasks {
		if v.Id == id {
			return &v
		}
	}

	return nil
}

func (r *inMemoryTaskRepository) getIndex(id int) int {
	for i, v := range r.tasks {
		if v.Id == id {
			return i
		}
	}

	return -1
}

func (r *inMemoryTaskRepository) remove(index int) []task.Task {
	return append(r.tasks[:index], r.tasks[index+1:]...)
}
