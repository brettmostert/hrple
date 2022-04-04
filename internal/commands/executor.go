package commands

import (
	"fmt"

	"github.com/brettmostert/hrple/pkg/cli"
)

type Executer struct {
	rootCommand *cli.Command
}

func NewExecuter() *Executer {
	rootCommand := &cli.Command{
		Name:             "hrple-cli",
		ShortDescription: "hrple-cli",
		LongDescription:  "Create a hrple of habits",
		Example:          "hrple-cli",
		Run:              DefaultCmd,
		// HelpType:         cli.Default, // Change this to be Help as Text then display that as default if nothing then generate
	}

	e := &Executer{
		rootCommand: rootCommand,
	}

	e.initTest()
	e.initAdd()

	return e
}

func (e *Executer) Execute() (interface{}, error) {
	return e.rootCommand.Execute()
}

func DefaultCmd(cmd *cli.Command, args []string) ([]interface{}, error) {
	fmt.Printf("Default Command")
	return nil, nil
}
