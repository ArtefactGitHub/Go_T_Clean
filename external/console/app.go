package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/interactor"
	"github.com/ArtefactGitHub/Go_T_Clean/external/console/command"
)

type App struct {
	interactor interactor.TaskInteractor
}

func (app *App) Run() {
	app.interactor = interactor.NewTaskInteractor()
	_ = app.stringPrompt("\n" + `Input <Command> or "help" or press Enter to exit.`)
}

func (app *App) stringPrompt(label string) error {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, label+"\n>>")

		s, _ = r.ReadString('\n')
		cmd, err := app.parseCommand(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n"+err.Error())
			continue
		}

		running, err := cmd.Do()
		if err != nil {
			fmt.Fprintln(os.Stderr, "\n"+err.Error())
		}
		if !running {
			fmt.Fprintln(os.Stdout, "exit")
			break
		}
	}
	return nil
}

func (app *App) parseCommand(input string) (command.Command, error) {
	trimed := strings.TrimSpace(input)

	splits := strings.Split(trimed, " ")
	if len(splits) == 1 {
		return command.NewCommand(splits[0], "", "", app.interactor)
	} else if len(splits) == 2 {
		return command.NewCommand(splits[0], splits[1], "", app.interactor)
	} else if len(splits) == 3 {
		return command.NewCommand(splits[0], splits[1], splits[2], app.interactor)
	}

	return nil, errors.New("invalid input. please check the help")
}
