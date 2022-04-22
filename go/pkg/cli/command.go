package cli

import (
	"os"
	"strings"

	"github.com/brettmostert/hrple/go/pkg/errors/exitError"
)

type Command struct {
	Name             string
	ShortDescription string
	LongDescription  string
	Example          string
	Run              func(cmd *Command, args []string) ([]interface{}, error)
	HelpType         Help
	Args             []string

	flagSet *FlagSet
	// flags   []*string

	parent   *Command
	commands []*Command
}

type Options struct {
	Args []string
}

func (cmd *Command) Execute(options ...Options) ([]interface{}, error) {
	args := os.Args[1:]

	if len(options) > 0 {
		args = options[0].Args
	}

	argsLen := len(args)

	var cmdToExecute *Command

	var argsToExecute []string

	switch {
	case argsLen == 0:
		cmdToExecute = cmd
		argsToExecute = args
	case argsLen > 0:
		cmdToExecute = cmd.findCommand(args)
		argsToExecute = args[1:]
	}

	if cmdToExecute == nil {
		return nil, exitError.New("Command not found, args: "+strings.Join(args, " "), exitError.NotFound)
	}

	err := cmdToExecute.parseFlags(argsToExecute)
	if err != nil {
		return nil, exitError.New("Unable to parse flags, args: "+strings.Join(args, " "), exitError.InvalidFlags)
	}

	return cmdToExecute.Run(cmdToExecute, argsToExecute)
}

func (cmd *Command) findNext(args []string) *Command {
	var nextCommand *Command

	for _, cmd := range cmd.commands {
		if cmd.Name == args[0] {
			nextCommand = cmd
			break
		}
	}

	return nextCommand
}

func (rootCmd *Command) findCommand(args []string) *Command {
	var command *Command

	var innerfind func(*Command, []string) *Command

	innerfind = func(innerCommand *Command, innerArgs []string) *Command {
		if len(innerArgs) == 0 {
			return innerCommand
		}

		command = innerCommand.findNext(innerArgs)
		if command != nil {
			command = innerfind(command, innerArgs[1:])
		}

		return command
	}

	parentCmd := rootCmd.findNext(args)
	if parentCmd != nil {
		argsWithoutFlags := stripFlags(args[1:])
		command = innerfind(parentCmd, argsWithoutFlags)

		if command == nil {
			if len(parentCmd.Args) > 0 {
				command = parentCmd
			}
		}
	}

	return command
}

func (parentCmd *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		if cmd == parentCmd {
			panic("Command cannot be a child of itself")
		}

		if parentCmd.Name == cmd.Name {
			panic("Parent command and child command cannot have the same name")
		}

		cmd.parent = parentCmd
		parentCmd.commands = append(parentCmd.commands, cmd)
	}
}
