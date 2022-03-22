package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ArtefactGitHub/Go_T_Clean/usecase/interfaces"
)

type delete struct {
	args []string
	id   int
	interfaces.TaskInteractor
}

func newDeleteCommand(args []string, intr interfaces.TaskInteractor) Command {
	cmd := delete{args: args, TaskInteractor: intr}
	return &cmd
}

func (cmd *delete) Do() (bool, error) {
	if len(cmd.args) != 2 {
		return true, errors.New("invalid argument")
	}

	if id, err := strconv.Atoi(cmd.args[1]); err != nil {
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
