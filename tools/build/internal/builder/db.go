package builder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/brettmostert/hrple/go/pkg/errors/exitError"
)

func (builder *Builder) Db(name string, ternArgs string) error {
	project, _ := builder.findProject(name)
	if project == nil {
		return exitError.New("Project not found, name: "+name, exitError.NotFound)
	}

	if _, err := os.Stat(project.Root + "/components/" + project.Name + "/db"); os.IsNotExist(err) {
		err := os.MkdirAll(project.Root+"/components/"+project.Name+"/db", 0700)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Executing: %v '%v'\n", "tern", ternArgs)

	cmd := exec.Command("tern", ternArgs)

	cmd.Dir = project.Root + "/components/" + project.Name + "/db"

	var out, outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", &outErr)
		return err
	}

	fmt.Printf("%s\n", &out)

	return nil
}
