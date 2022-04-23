package builder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type BuildConfig struct {
	Config   Config    `json:"config"`
	Projects []Project `json:"projects"`
}

type Config struct {
	Output string `json:"output"`
}

type Project struct {
	Name     string    `json:"name"`
	Language string    `json:"lang"`
	Type     string    `json:"type"`
	Path     string    `json:"path"`
	Root     string    `json:"root"`
	Releases []Release `json:releases`
}

type Release struct {
	Name    string   `json:"name"`
	Flags   []string `json:"flags"`
	Default bool     `json:"default"`
}

func readFile(jsonFilePath string) ([]byte, error) {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}

func parseConfig(configJson []byte) (*BuildConfig, error) {
	var buildConfig BuildConfig

	err := json.Unmarshal(configJson, &buildConfig)
	if err != nil {
		return nil, err
	}

	return &buildConfig, nil
}

func Print(jsonFilePath string) {
	bytes, _ := readFile(jsonFilePath)

	buildConfig, _ := parseConfig(bytes)

	b, err := json.MarshalIndent(buildConfig, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(b))
}
