package commands

import (
	"fmt"

	"github.com/brettmostert/hrple/go/pkg/cli"
)

func (e *Executer) initAdd() {
	cmd := &cli.Command{
		Name: "add",
		Run:  ExecuteAdd,
		// Args: []string{"name"},
	}

	cmd.Args().Set("name")

	e.rootCommand.AddCommand(cmd)
}

func ExecuteAdd(cmd *cli.Command, args []string) ([]interface{}, error) {
	fmt.Printf("ExecuteAdd with these args: %v", args)
	return nil, nil
}
