package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/brettmostert/hrple/go/pkg/errors/exitError"
)

type Builder struct {
	buildConfig *BuildConfig
	filePath    string
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
		filePath:    filePath,
	}

	return builder
}

func (builder *Builder) UpdateConfigFile(data []byte) error {
	return ioutil.WriteFile(builder.filePath, data, fs.ModeAppend)
}

func (builder *Builder) AddProject(project *Project) error {
	existingProject := builder.findProject(project.Name)
	if existingProject != nil {
		return exitError.New("Project already exists, name: ", exitError.Failure)
	}

	if project.Root == "" {
		project.Root = project.Language
	}

	if project.Path == "" {
		project.Path = "./components/" + project.Name
		if project.Language == "go" {
			project.Path = project.Path + "/cmd/" + project.Name
		}
	}

	project.Releases = append(project.Releases, Release{
		Name:    "default",
		Default: true,
		Flags:   []string{},
	})

	builder.buildConfig.Projects = append(builder.buildConfig.Projects, *project)

	// TODO: Move this to create for other languages & to create other folders and potentially "main" file i.e main.go
	err := os.MkdirAll("./"+project.Language+"/"+project.Path[1:], 0700)
	if err != nil {
		return err
	}

	data, _ := json.MarshalIndent(builder.buildConfig, "", "\t")
	return builder.UpdateConfigFile(data)

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
		if rel.Default {
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
