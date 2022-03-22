package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/ArtefactGitHub/Go_T_Clean/usecase/interfaces"
)

type readall struct {
	args []string
	interfaces.TaskInteractor
}

func newReadAllCommand(args []string, intr interfaces.TaskInteractor) Command {
	cmd := readall{args: args, TaskInteractor: intr}
	return &cmd
}

func (cmd *readall) Do() (bool, error) {
	if len(cmd.args) != 1 {
		return true, errors.New("invalid argument")
	}

	tasks, err := cmd.TaskInteractor.GetAll()
	if err != nil {
		return false, err
	}

	result := ""
	for _, v := range tasks {
		result += fmt.Sprintf("[%d][%s]\n", v.Id, v.Name)
	}
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
