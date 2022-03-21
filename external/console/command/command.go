package command

import (
	"github.com/ArtefactGitHub/Go_T_Clean/external/console/constant"
)

type CommandParm struct {
	CommandType constant.CommandType
	Arg         interface{}
}

func NewCommandParam(commandName string, arg interface{}) CommandParm {
	switch commandName {
	case "readall":
		return CommandParm{CommandType: constant.ReadAll, Arg: arg}
	case "read":
		return CommandParm{CommandType: constant.Read, Arg: arg}
	case "create":
		return CommandParm{CommandType: constant.Create, Arg: arg}
	case "update":
		return CommandParm{CommandType: constant.Update, Arg: arg}
	case "delete":
		return CommandParm{CommandType: constant.Delete, Arg: arg}
	case "help":
		return CommandParm{CommandType: constant.Help, Arg: arg}
	case "exit", "":
		return CommandParm{CommandType: constant.Exit, Arg: arg}
	}

	return CommandParm{CommandType: constant.None, Arg: arg}
}
