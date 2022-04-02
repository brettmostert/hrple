package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/brettmostert/hrple/pkg/errors/exitError"
)

type Command struct {
	Name             string
	ShortDescription string
	LongDescription  string
	Example          string
	Run              func(cmd *Command, args []string) error
	HelpType         Help
	Args             []string

	flagSet *flag.FlagSet
	flags   []*string

	parent   *Command
	commands []*Command
}

func (parentCmd *Command) Execute() error {
	args := os.Args[1:]
	argsLen := len(args)

	if argsLen == 0 {
		return parentCmd.Run(parentCmd, args)
	}

	var cmdToExecute *Command

	switch {
	case argsLen == 0:
		cmdToExecute = parentCmd
	case argsLen > 0:
		cmdToExecute = parentCmd.findSubCommandToExecute(args)
	}

	if cmdToExecute == nil {
		return exitError.New("Command not found", exitError.NotFound)
	}

	return cmdToExecute.Run(parentCmd, args[1:])
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

func (rootCmd *Command) findSubCommandToExecute(args []string) *Command {
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

	// flag.Args()
	parentCmd := rootCmd.findNext(args)
	if parentCmd != nil {
		fmt.Printf("Logging, parentCmd->Command.Name: %v \n", parentCmd.Name)

		// TODO: After implementing UnitTests, clean up, implement flags and args
		// command = innerfind(command, args[1:])
		// a := args[1:]
		// if command.flagSet != nil {
		// 	a = command.flagSet.Args()
		// }

		command = innerfind(parentCmd, args[1:])
		if command != nil {
			fmt.Printf("Logging, childCmd->Command.Name: %v \n", command.Name)
		} else {
			fmt.Printf("Logging, Command is Nil, Parent Command %v\n", parentCmd.Name)
			if len(parentCmd.Args) > 0 {
				command = parentCmd
			}
		}
	}

	// TODO: After implementing UnitTests, clean up, implement flags and args
	// l := command.flags[0]
	// command.flagSet.Parse(args[1:])
	// fmt.Printf("flags Name: %v FlagSet: %v Flags: %v \n", command.Name, command.flagSet.Args(), *l)

	// setup a flag check function, if no flag set creat one, and return flagSet (see flags() from
	//	https://github.com/spf13/cobra/blob/cb9d7b1cec87c2bb005c6e2790553bcd629bc542/command.go#L1450
	// call in the innerFind
	// do we need flags? or should I simply use the next values in order...? check out cmd line options of cobra
	// ie cli create X -p parent

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

func (cmd *Command) AddFlag() {
	if cmd.flagSet == nil {
		cmd.flagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	}

	cmd.flags = append(cmd.flags, cmd.flagSet.String("v", "verbose", ""))
}
