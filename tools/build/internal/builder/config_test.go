package builder

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	filePath := "../../testing/build.json"
	
	_, err := readFile(filePath)

	if (err != nil) {
		t.Errorf("Did not expect error '%v'", err)
	}
}


func TestParseConfig(t *testing.T) {
	filePath := "../../testing/build.json"
	
	bytes, _ := readFile(filePath)

	buildConfig, err := parseConfig(bytes)

	if (err != nil) {
		t.Errorf("Did not expect error '%v'", err)
	}

	if buildConfig.Config.Output != "dist" {
		t.Errorf("Expected '%v', got: %v", "dist", buildConfig.Config.Output)
	}

	if !(len(buildConfig.Projects) > 0) {
		t.Errorf("Expected len of Project to be > '%v', got: %v", 0, len(buildConfig.Projects))
	}

	if buildConfig.Projects[0].Name != "hrple-cli" {
		t.Errorf("Expected '%v', got: %v", "hrple-cli", buildConfig.Projects[0].Name)
	}

	fmt.Printf("buildConfig : %v", buildConfig.Config.Output)
}