package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/domain/model"
)

type update struct {
	idStr string
	name  string
	id    int
	interactor.TaskInteractor
}

func newUpdateCommand(idStr string, name string, intr interactor.TaskInteractor) Command {
	cmd := update{idStr: idStr, name: name, TaskInteractor: intr}
	return &cmd
}

func (cmd *update) Do() (bool, error) {
	if id, err := strconv.Atoi(cmd.idStr); err != nil {
		return true, errors.New("invalid argument")
	} else {
		cmd.id = id
	}

	task := model.NewTask(cmd.id, cmd.name)
	updated, err := cmd.TaskInteractor.Update(task)
	if err != nil {
		return false, err
	}

	result := fmt.Sprintf("update success. [%d][%s]\n", updated.Id, updated.Name)
	fmt.Fprint(os.Stdout, result)
	return true, nil
}
