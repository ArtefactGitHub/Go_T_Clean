package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ArtefactGitHub/Go_T_Clean/domain/model/task"
	"github.com/ArtefactGitHub/Go_T_Clean/external/common"
	"github.com/ArtefactGitHub/Go_T_Clean/external/console/command"
	itask "github.com/ArtefactGitHub/Go_T_Clean/external/infurastructure/persistence/inmemory/task"
	utask "github.com/ArtefactGitHub/Go_T_Clean/usecase/task"
)

type consoleApp struct {
	deployType common.DeployType
	storeType  common.StoreType
	interactor utask.TaskInteractor
	repository task.TaskRepository
}

func NewConsoleApp(deployType common.DeployType, storeType common.StoreType) common.App {
	app := consoleApp{deployType: deployType, storeType: storeType}
	return &app
}

func (app *consoleApp) Run() error {
	var err error
	app.repository, err = itask.NewInMemoryTaskRepository()
	if err != nil {
		return err
	}
	app.interactor = utask.NewTaskInteractor(app.repository)

	return app.stringPrompt("\n" + `Input <Command> or "help" or press Enter to exit.`)
}

func (app *consoleApp) stringPrompt(label string) error {
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

func (app *consoleApp) parseCommand(input string) (command.Command, error) {
	trimed := strings.TrimSpace(input)

	splits := strings.Split(trimed, " ")
	return command.NewCommand(app.interactor, splits)
}
