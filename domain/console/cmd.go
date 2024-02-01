package console

import "github.com/spf13/cobra"

type ScheduleFunc func(cmd *cobra.Command) []string
type ScheduleMap map[string]ScheduleFunc
type ScheduleType map[string]ScheduleMap

type SetCommandInterface interface {
	PutCommand()
	GetCommand()
}

var command *Command

type Command struct {
	Cmd *cobra.Command
}

func NewConsole() *Command {
	command = &Command{}
	return command
}

func GetCmd() *cobra.Command {
	return command.Cmd
}

func (cmd *Command) PutCommand(operateCmd *cobra.Command) {
	cmd.Cmd = operateCmd
}

func (cmd *Command) GetCommand() *cobra.Command {
	return cmd.Cmd
}
