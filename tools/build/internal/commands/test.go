package commands

import (
	"github.com/brettmostert/hrple/go/pkg/cli"
	config "github.com/brettmostert/hrple/tools/build/internal/builder"
)

func (e *Executer) initTest() {
	cmd := &cli.Command{
		Name: "test",
		Run:  ExecuteTest,
		Args: 			[]string{"project"},
	}	

	cmd.Flags().String("f", "build.json", "")

	e.rootCommand.AddCommand(cmd)
}

func ExecuteTest(cmd *cli.Command, args []string) ([]interface{}, error) {
	builder := config.NewBuilder(cmd.Flags().GetString("f"))
	_ = builder.Test(args[0])
	
	return nil, nil
}