package cli

import (
	"flag"
	"strconv"
	"strings"
)

// A wrapper to extend functionality.
type FlagSet struct {
	fs *flag.FlagSet
}

func (cmd *Command) Flags() *FlagSet {
	if cmd.flagSet == nil {
		cmd.flagSet = &FlagSet{
			fs: flag.NewFlagSet(cmd.Name, flag.ContinueOnError),
		}
	}

	return cmd.flagSet
}

func (cmd *Command) parseFlags(args []string) error {
	flags := []string{}
	numberOfArgs := len(args) - 1

	for index, arg := range args {
		nextArg := ""

		if numberOfArgs != index {
			nextArg = args[index+1]
		}

		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			flags = append(flags, arg)
			// add value for - or --arg moo.
			if nextArg != "" && !strings.HasPrefix(nextArg, "-") && !strings.HasPrefix(nextArg, "--") {
				flags = append(flags, nextArg)
			}
		}
	}

	return cmd.Flags().fs.Parse(flags)
}

func stripFlags(args []string) []string {
	argsWithoutFlags := []string{}

	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			argsWithoutFlags = append(argsWithoutFlags, arg)
		}
	}

	return argsWithoutFlags
}

func (flagSet *FlagSet) GetBool(name string) (bool, error) {
	return strconv.ParseBool(flagSet.fs.Lookup(name).Value.String())
}

func (flagSet *FlagSet) Bool(name string, defaultValue bool, usage string) *bool {
	return flagSet.fs.Bool(name, defaultValue, usage)
}

func (flagSet *FlagSet) Value(name string) string {
	return flagSet.fs.Lookup(name).Value.String()
}

func (flagSet *FlagSet) GetString(name string) string {
	return flagSet.Value(name)
}

func (flagSet *FlagSet) String(name string, defaultValue string, usage string) *string {
	return flagSet.fs.String(name, defaultValue, usage)
}

func (flagSet *FlagSet) Int(name string, defaultValue int, usage string) *int {
	return flagSet.fs.Int(name, defaultValue, usage)
}

func (flagSet *FlagSet) GetInt(name string) (int, error) {
	return strconv.Atoi(flagSet.fs.Lookup(name).Value.String())
}
