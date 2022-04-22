package commands

import (
	"fmt"

	"github.com/brettmostert/hrple/go/pkg/cli"
)

func (e *Executer) initTest() {
	cmd := &cli.Command{
		Name: "test",
		Run:  ExecuteTest,
		Args: []string{"moo"},
	}

	cmd.Flags().String("name", "", "help message for name")
	cmd.Flags().String("lastname", "meow", "help message for lastname")
	cmd.Flags().Bool("verbose", false, "help message for verbose")
	cmd.Flags().Bool("party", true, "help message for party")

	subCmd := &cli.Command{
		Name: "a",
		Run:  ExecuteA,
	}

	cmd.AddCommand(subCmd)
	e.rootCommand.AddCommand(cmd)
}

func ExecuteTest(cmd *cli.Command, args []string) ([]interface{}, error) {
	verboseFlagValue, err := cmd.Flags().GetBool("verbose")
	fmt.Printf("ExecuteTest %v %v %v \n", verboseFlagValue, err, cmd.Flags().GetString("name"))

	return nil, nil
}

func ExecuteA(cmd *cli.Command, args []string) ([]interface{}, error) {
	// return nil, exitError.New("ExecuteA, not implemented", exitError.NotImplemented)
	fmt.Printf("ExecuteA %v", args)
	fmt.Printf("ExecuteA %v", cmd.Flags().GetString("name"))

	return nil, nil
}
