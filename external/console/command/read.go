package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
)

type read struct {
	idStr string
	id    int
	interactor.TaskInteractor
}

func newReadCommand(idStr string, intr interactor.TaskInteractor) Command {
	cmd := read{idStr: idStr, TaskInteractor: intr}
	return &cmd
}

func (cmd *read) Do() (bool, error) {
	if id, err := strconv.Atoi(cmd.idStr); err != nil {
		return true, errors.New("invalid argument")
	} else {
		cmd.id = id
	}

	task, err := cmd.TaskInteractor.Get(cmd.id)
	if err != nil {
		return false, err
	} else if task == nil {
		return true, fmt.Errorf("not found. id: %d", cmd.id)
	}

	result := fmt.Sprintf("[%d][%s]\n", task.Id, task.Name)
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
