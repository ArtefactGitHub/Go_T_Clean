package command

import (
	"fmt"
	"os"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
)

type create struct {
	name string
	interactor.TaskInteractor
}

func newCreateCommand(name string, intr interactor.TaskInteractor) Command {
	cmd := create{name: name, TaskInteractor: intr}
	return &cmd
}

func (cmd *create) Do() (bool, error) {
	task := model.NewTask(0, cmd.name)
	id, err := cmd.TaskInteractor.Create(task)
	if err != nil {
		return false, err
	}

	result := fmt.Sprintf("create success. [%d][%s]\n", id, task.Name)
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
