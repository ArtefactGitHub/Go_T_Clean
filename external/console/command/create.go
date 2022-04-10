package command

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model/task"
	utask "github.com/ArtefactGitHub/Go_T_Clean/usecase/task"
)

type create struct {
	args []string
	utask.TaskInteractor
}

func newCreateCommand(args []string, intr utask.TaskInteractor) Command {
	cmd := create{args: args, TaskInteractor: intr}
	return &cmd
}

func (cmd *create) Do() (bool, error) {
	if len(cmd.args) != 2 {
		return true, errors.New("invalid argument")
	}

	task := task.NewTask(0, cmd.args[1])
	id, err := cmd.TaskInteractor.Create(context.TODO(), task)
	if err != nil {
		return false, err
	}

	result := fmt.Sprintf("create success.\n[%d][%s]\n", id, task.Name)
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
