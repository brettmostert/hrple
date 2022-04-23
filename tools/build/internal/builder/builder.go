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
	if err != nil {
		panic(err)
	}

	buildConfig, err := parseConfig(b)
	if err != nil {
		panic(err)
	}

	builder := &Builder{
		buildConfig: buildConfig,
	}

	return builder
}

func (builder *Builder) Build(name string, releaseName string) error {
	project := builder.findProject(name)
	if project == nil {
		return exitError.New("Project not found, name: "+name, exitError.NotFound)
	}

	release := project.findRelease(releaseName)
	if release == nil {
		release = project.findDefault(releaseName)
		if release == nil {
			return exitError.New("Release not found, name: "+releaseName, exitError.NotFound)
		}
	}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	output := []string{path, "/", builder.buildConfig.Config.Output, "/", project.Name, "/", release.Name, "/"}

	args := []string{"build", "-o", strings.Join(output, "")}

	if len(release.Flags) > 0 {
		args = append(args, release.Flags...)
	}

	args = append(args, project.Path)

	fmt.Printf("Executing: %v %v", "go", strings.Join(args, " "))
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

func (project *Project) findRelease(name string) *Release {
	for _, rel := range project.Releases {
		if strings.EqualFold(rel.Name, name) {
			return &rel
		}
	}

	return nil
}

func (project *Project) findDefault(name string) *Release {
	for _, rel := range project.Releases {
		if rel.Default == true {
			return &rel
		}
	}

	return nil
}

func (builder *Builder) Test(name string) error {
	project := builder.findProject(name)
	if project == nil {
		return exitError.New("Project not found, name: "+name, exitError.NotFound)
	}

	args := []string{"test", "./...", "-cover", "-covermode=atomic"}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	runDir := path + "/" + project.Root

	// Bob is testing him self again.
	if strings.Contains(path, "tools/build/internal/builder") {
		runDir = strings.Replace(path, "tools/build/internal/builder", project.Root, 1)
	}

	cmd := exec.Command("go", args...)
	cmd.Dir = runDir

	var out, outErr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &outErr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s \n %s \n", &out, &outErr)
		return err
	}

	fmt.Printf("%s\n", &out)

	return nil
}
