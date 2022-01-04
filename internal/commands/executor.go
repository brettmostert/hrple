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
		Run:              ExecuteCmd,
		HelpType:         cli.Default,
	}

	e := &Executer{
		rootCommand: rootCommand,
	}

	e.initTest()
	e.initAdd()

	return e
}

func (e *Executer) Execute() error {
	return e.rootCommand.Execute()
}

func ExecuteCmd(cmd *cli.Command, args []string) error {
	fmt.Printf("root")
	return nil
}
