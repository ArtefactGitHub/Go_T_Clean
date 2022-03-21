package command

import (
	"errors"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/external/console/constant"
)

type Command interface {
	Do() (bool, error)
}

func NewCommand(intr interactor.TaskInteractor, args []string) (Command, error) {
	length := len(args)
	if length == 0 {
		return nil, errors.New("invalid input. please check the help")
	}

	cmdType, err := constant.ParseCommandType(args[0])
	if err != nil {
		return nil, err
	}

	switch cmdType {
	case constant.ReadAll:
		return newReadAllCommand(args, intr), nil
	case constant.Read:
		return newReadCommand(args, intr), nil
	case constant.Create:
		return newCreateCommand(args, intr), nil
	case constant.Update:
		return newUpdateCommand(args, intr), nil
	case constant.Delete:
		return newDeleteCommand(args, intr), nil
	case constant.Help:
		return newHelpCommand(), nil
	case constant.Exit:
		return newExitCommand(), nil
	}

	return newNoneCommand(), nil
}
