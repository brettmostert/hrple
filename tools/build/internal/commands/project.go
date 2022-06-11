package commands

import (
	"fmt"

	"github.com/brettmostert/hrple/go/pkg/cli"
	config "github.com/brettmostert/hrple/tools/build/internal/builder"
)

func (e *Executer) initProject() {
	cmd := &cli.Command{
		Name: "project",
		Run:  ExecuteProject,
	}

	cmdAdd := cli.Command{
		Name: "add",
		Run:  ExecuteAddProject,
		// Args: []string{"project"},
	}
	cmdAdd.Args().Set("project")
	cmdAdd.Flags().String("f", "build.json", "")
	cmdAdd.Flags().String("lang", "go", "")
	cmdAdd.Flags().String("type", "component", "")
	cmdAdd.Flags().String("root", "", "") // default is language name
	cmdAdd.Flags().String("path", "", "") // default {lang}/components/{project}/cmd/{name} for go

	cmd.AddCommand(&cmdAdd)

	cmdRemove := cli.Command{
		Name: "remove",
		Run:  ExecuteRemoveProject,
		// Args: []string{"project"},
	}
	cmdRemove.Args().Set("project")
	cmdRemove.Flags().String("f", "build.json", "")

	cmdRm := cli.Command{
		Name: "rm",
		Run:  ExecuteRemoveProject,
		// Args: []string{"project"},
	}

	cmdRm.Args().Set("project")
	cmdRm.Flags().String("f", "build.json", "")
	cmd.AddCommand(&cmdRemove)
	cmd.AddCommand(&cmdRm)

	e.rootCommand.AddCommand(cmd)
}

func ExecuteProject(cmd *cli.Command, args []string) ([]interface{}, error) {
	fmt.Println("project222")

	return nil, nil
}

func ExecuteAddProject(cmd *cli.Command, args []string) ([]interface{}, error) {
	builder := config.NewBuilder(cmd.Flags().GetString("f"))

	project := &config.Project{
		Name:     cmd.Args().Get("project"),
		Language: cmd.Flags().GetString("lang"),
		Type:     cmd.Flags().GetString("type"),
		Path:     cmd.Flags().GetString("path"),
		Root:     cmd.Flags().GetString("root"),
	}

	err := builder.AddProject(project)

	return nil, err
}

func ExecuteRemoveProject(cmd *cli.Command, args []string) ([]interface{}, error) {
	builder := config.NewBuilder(cmd.Flags().GetString("f"))

	err := builder.RemoveProject(cmd.Args().Get("project"))

	return nil, err
}
