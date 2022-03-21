package command

import (
	"fmt"
	"os"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
)

type readall struct {
	interactor.TaskInteractor
}

func newReadAllCommand(intr interactor.TaskInteractor) Command {
	cmd := readall{TaskInteractor: intr}
	return &cmd
}

func (cmd *readall) Do() (bool, error) {
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
