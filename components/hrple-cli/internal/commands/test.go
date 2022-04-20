package commands

import (
	"github.com/brettmostert/hrple/pkg/cli"
	"github.com/brettmostert/hrple/pkg/errors/exitError"
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
	return nil, exitError.New("ExecuteTest, not implemented", exitError.NotImplemented)
}

func ExecuteA(cmd *cli.Command, args []string) ([]interface{}, error) {
	return nil, exitError.New("ExecuteA, not implemented", exitError.NotImplemented)
}
