package command

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model/task"
	utask "github.com/ArtefactGitHub/Go_T_Clean/usecase/task"
)

type update struct {
	args []string
	id   int
	utask.TaskInteractor
}

func newUpdateCommand(args []string, intr utask.TaskInteractor) Command {
	cmd := update{args: args, TaskInteractor: intr}
	return &cmd
}

func (cmd *update) Do() (bool, error) {
	if len(cmd.args) != 3 {
		return true, errors.New("invalid argument")
	}
	fmt.Fprintf(os.Stdout, "%v\n", cmd.args)

	if id, err := strconv.Atoi(cmd.args[1]); err != nil {
		return true, errors.New("invalid argument")
	} else {
		cmd.id = id
	}

	task := task.NewTask(cmd.id, cmd.args[2])
	updated, err := cmd.TaskInteractor.Update(context.TODO(), task)
	if err != nil {
		return false, err
	} else if updated == nil {
		return true, fmt.Errorf("not found. id: %d", cmd.id)
	}

	result := fmt.Sprintf("update success.\n[%d][%s]\n", updated.Id, updated.Name)
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
