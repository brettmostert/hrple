package commands

import (
	"fmt"

	"github.com/brettmostert/hrple/go/pkg/cli"
)

type Executer struct {
	rootCommand *cli.Command
}

func NewExecuter() *Executer {
	rootCommand := &cli.Command{
		Name:             "bob",
		ShortDescription: "Bob the Builder",
		LongDescription:  "A builder builds what a builder builds",
		Example:          "bob build {projectName}",
		Run:              DefaultCmd,
	}

	e := &Executer{
		rootCommand: rootCommand,
	}
	
	e.initConfig()
	e.initBuild()
	e.initTest()

	return e
}

func (e *Executer) Execute() (interface{}, error) {
	return e.rootCommand.Execute()
}

func DefaultCmd(cmd *cli.Command, args []string) ([]interface{}, error) {
	fmt.Printf("buildesr %v", args)
	return nil, nil
}
