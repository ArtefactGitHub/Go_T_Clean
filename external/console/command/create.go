package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
)

type create struct {
	args []string
	interactor.TaskInteractor
}

func newCreateCommand(args []string, intr interactor.TaskInteractor) Command {
	cmd := create{args: args, TaskInteractor: intr}
	return &cmd
}

func (cmd *create) Do() (bool, error) {
	if len(cmd.args) != 2 {
		return true, errors.New("invalid argument")
	}

	task := model.NewTask(0, cmd.args[1])
	id, err := cmd.TaskInteractor.Create(task)
	if err != nil {
		return false, err
	}

	result := fmt.Sprintf("create success. [%d][%s]\n", id, task.Name)
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
