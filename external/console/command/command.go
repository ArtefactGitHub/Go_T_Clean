package command

import (
	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/external/console/constant"
)

type Command interface {
	Do() (bool, error)
}

func NewCommand(commandName string, arg string, arg2 string, intr interactor.TaskInteractor) (Command, error) {
	cmdType, err := constant.ParseCommandType(commandName)
	if err != nil {
		return nil, err
	}

	switch cmdType {
	case constant.ReadAll:
		return newReadAllCommand(intr), nil
	case constant.Read:
		return newReadCommand(arg, intr), nil
	case constant.Create:
		return newCreateCommand(arg, intr), nil
	case constant.Update:
		return newUpdateCommand(arg, arg2, intr), nil
	case constant.Delete:
		return newDeleteCommand(arg, intr), nil
	case constant.Help:
		return newHelpCommand(), nil
	case constant.Exit, constant.None:
		return newExitCommand(), nil
	}

	return newExitCommand(), nil
}
