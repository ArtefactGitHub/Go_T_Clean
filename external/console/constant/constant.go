package constant

import "fmt"

const (
	HelpMessage = `
Read All Task : "readall"
Read Task     : "read <task id>"
Create Task   : "create <task name>"
Update Task   : "update <task id> <task name>"
Delete Task   : "delete <task id>"`
)

type CommandType string

const (
	ReadAll CommandType = "readall"
	Read    CommandType = "read"
	Create  CommandType = "create"
	Update  CommandType = "update"
	Delete  CommandType = "delete"
	Help    CommandType = "help"
	Exit    CommandType = "exit"
	None    CommandType = ""
)

func ParseCommandType(s string) (v CommandType, err error) {
	v = CommandType(s)
	err = v.valid()
	return v, err
}

func (v CommandType) valid() error {
	switch v {
	case ReadAll, Read, Create, Update, Delete, Help, Exit, None:
		return nil
	default:
		return fmt.Errorf("invalid CommandType: %s", v)
	}
}
