package builder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/brettmostert/hrple/go/pkg/errors/exitError"
)

type Builder struct {
	buildConfig *BuildConfig
}

func NewBuilder(filePath string) *Builder {
	b, err := readFile(filePath)
	if (err != nil) {
		panic(err)
	}
	
	buildConfig, err := parseConfig(b)
	if (err != nil) {
		panic(err)
	}
	
	builder := &Builder{
		buildConfig: buildConfig,
	}
	
	return builder
}

func (builder *Builder) Build(name string) error {
	project := builder.findProject(name)
	if (project == nil) {
		return exitError.New("Project not found, name: " + name, exitError.NotFound)
	}
	
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	output := []string{path,"/",builder.buildConfig.Config.Output,"/", project.Name,"/"}
	
	args := []string{"build","-o", strings.Join(output,""), project.Path }
	
	cmd := exec.Command("go", args...)
	cmd.Dir = project.Root

	var out, outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", &outErr)
		return err
	}
	
	fmt.Printf("%s\n", &out)

	return nil
}

func (builder *Builder) findProject(name string) *Project {
	for _, project := range builder.buildConfig.Projects {
		if strings.EqualFold(project.Name, name) {			
			return &project
		}		
	}

	return nil
}

func (builder *Builder) Test(name string) error {
	project := builder.findProject(name)
	if (project == nil) {
		return exitError.New("Project not found, name: " + name, exitError.NotFound)
	}
	
	args := []string{"test", "./...", "-cover", "-covermode=atomic"}
	
	cmd := exec.Command("go", args...)
	cmd.Dir = project.Root

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