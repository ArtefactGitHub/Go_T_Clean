package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ArtefactGitHub/Go_T_Clean/external/console/command"
	"github.com/ArtefactGitHub/Go_T_Clean/external/console/constant"
)

type App struct{}

const (
	mark = ">>"
	help = `
Read All Task : "readall"
Read Task     : "read <task id>"
Create Task   : "create <task name>"
Update Task   : "update <task id> <task name>"
Delete Task   : "delete <task id>"`
)

func (app *App) Run() {
	_ = StringPrompt("\n" + `Input <Command> or "help" or press Enter to exit.`)
}

func StringPrompt(label string) error {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, label+"\n"+mark)

		s, _ = r.ReadString('\n')
		cmd, err := ParseCommand(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n"+err.Error())
			break
		}

		running, err := DoCommand(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n"+err.Error())
			break
		} else if !running {
			fmt.Fprintln(os.Stdout, "exit")
			break
		}
	}
	return nil
}

func ParseCommand(input string) (command.CommandParm, error) {
	trimed := strings.TrimSpace(input)
	fmt.Fprintf(os.Stdout, "command: [%s]\n", trimed)

	splits := strings.Split(trimed, " ")
	if len(splits) == 1 {
		return command.NewCommandParam(splits[0], nil), nil
	} else if len(splits) == 2 {
		return command.NewCommandParam(splits[0], splits[1]), nil
	}

	return command.NewCommandParam("", nil), errors.New("invalid input. please check the help")
}

func DoCommand(command command.CommandParm) (bool, error) {
	switch command.CommandType {
	case constant.ReadAll:
	case constant.Read:
	case constant.Create:
	case constant.Update:
	case constant.Delete:
	case constant.Help:
		fmt.Fprintln(os.Stdout, help)
		return true, nil
	case constant.Exit:
		return false, nil
	default:
	}
	return true, nil
}
