package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type App struct{}

const (
	mark = ">>"
	help = `
Create Task   : "create <task name>"
Read All Task : "read all"
Read Task     : "read <task id>"
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
		_, err := ParseCommand(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n"+err.Error())
			break
		}
	}
	return nil
}

type CommandParm struct{}

func ParseCommand(input string) (*CommandParm, error) {
	if input == "\n" {
		return nil, errors.New("exit")
	}
	if input == "help\n" {
		fmt.Fprintln(os.Stdout, help)
		return nil, nil
	}

	trimed := strings.TrimSpace(input)
	fmt.Fprintf(os.Stdout, "command: [%s]\n", trimed)

	return nil, nil
}
