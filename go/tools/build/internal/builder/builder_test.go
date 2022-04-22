package builder

import (
	"strings"
	"testing"
)

func TestFindProject(t *testing.T) {
	filePath := "../../testing/build.json"
	builder := NewBuilder(filePath)

	project := builder.findProject("hrple-cli")

	if (project == nil) {
		t.Errorf("Expected to find project, got: nil")	
		return
	}

	if !strings.EqualFold(project.Language, "gO") {
		t.Errorf("Expected language to be GO")	
	}
}

func TestFindProjectDoesNotExist(t *testing.T) {
	filePath := "../../testing/build.json"
	builder := NewBuilder(filePath)

	project := builder.findProject("404")

	if (project != nil) {
		t.Errorf("Expected NOT to find project")	
	}
}

func TestTest(t *testing.T) {
	filePath := "../../testing/build.json"
	builder := NewBuilder(filePath)

	err := builder.Test("hrple-cli")

	if (err != nil) {
		t.Errorf("Expected no errors, got: %v", err)	
	}
}

func TestTestNotFound(t *testing.T) {
	filePath := "../../testing/build.json"
	builder := NewBuilder(filePath)

	err := builder.Test("404")

	if (err.Error() != "Project not found, name: 404") {
		t.Errorf("Expected no errors, got: %v", err)	
	}
}


func TestBuild(t *testing.T) {
	t.Errorf("Not Implemented")
}

