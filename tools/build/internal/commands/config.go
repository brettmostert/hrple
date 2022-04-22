package commands

import (
	"github.com/brettmostert/hrple/go/pkg/cli"
	config "github.com/brettmostert/hrple/tools/build/internal/builder"
)

func (e *Executer) initConfig() {
	cmd := &cli.Command{
		Name: "config",
		Run:  ExecuteConfig,
		// Args: []string{"moo"},
	}	

	subCmd := &cli.Command{
		Name: "print",
		Run:  ExecutePrint,
	}

	subCmd.Flags().String("f", "build.json", "")

	cmd.AddCommand(subCmd)
	e.rootCommand.AddCommand(cmd)
}

func ExecuteConfig(cmd *cli.Command, args []string) ([]interface{}, error) {
	return nil, nil
}

func ExecutePrint(cmd *cli.Command, args []string) ([]interface{}, error) {
	config.Print(cmd.Flags().GetString("f"))
	return nil, nil
}
