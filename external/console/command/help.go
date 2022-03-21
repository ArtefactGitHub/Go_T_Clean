package command

import (
	"fmt"
	"os"

	"github.com/ArtefactGitHub/Go_T_Clean/external/console/constant"
)

type help struct {
}

func newHelpCommand() Command {
	cmd := help{}
	return &cmd
}

func (cmd *help) Do() (bool, error) {
	fmt.Fprintln(os.Stdout, constant.HelpMessage)
	return true, nil
}
