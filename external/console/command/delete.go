package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
)

type delete struct {
	idStr string
	id    int
	interactor.TaskInteractor
}

func newDeleteCommand(idStr string, intr interactor.TaskInteractor) Command {
	cmd := delete{idStr: idStr, TaskInteractor: intr}
	return &cmd
}

func (cmd *delete) Do() (bool, error) {
	if id, err := strconv.Atoi(cmd.idStr); err != nil {
		return true, errors.New("invalid argument")
	} else {
		cmd.id = id
	}

	success, err := cmd.TaskInteractor.Delete(cmd.id)
	if err != nil {
		return false, err
	}

	result := fmt.Sprintf("delete success: %v\n", success)
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
