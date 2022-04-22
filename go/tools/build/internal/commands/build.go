package commands

import (
	"github.com/brettmostert/hrple/go/pkg/cli"
	config "github.com/brettmostert/hrple/go/tools/build/internal/builder"
)

func (e *Executer) initBuild() {
	cmd := &cli.Command{
		Name: "build",
		Run:  ExecuteBuild,
		Args: 			[]string{"project"},
	}	

	cmd.Flags().String("f", "build.json", "")

	e.rootCommand.AddCommand(cmd)
}

func ExecuteBuild(cmd *cli.Command, args []string) ([]interface{}, error) {
	builder := config.NewBuilder(cmd.Flags().GetString("f"))
	
	_ = builder.Build(args[0])

	return nil, nil
}