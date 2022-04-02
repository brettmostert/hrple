package cli

import (
	"fmt"
	"strings"
	"testing"
)

func SetupCliForTesting() *Command {
	rootCommand := &Command{
		Name: "cli",
	}

	rootCommand.AddCommand(&Command{
		Name: "subCmdA",
	})

	subCmdB := &Command{
		Name: "subCmdB",
	}

	subCmdB.AddCommand(&Command{
		Name: "subCmdB-1",
	})

	rootCommand.AddCommand(subCmdB)

	subCmdC := &Command{
		Name: "subCmdC",
		Args: []string{"arg1", "arg2"},
	}

	rootCommand.AddCommand(subCmdC)

	return rootCommand
}

func TestFindCommandToExecute(t *testing.T) {
	rootCommand := SetupCliForTesting()

	var tests = []struct {
		name     string
		command  string
		expected string
	}{
		{"find subCmdA", "subCmdA", "subCmdA"},
		{"find subCmdB", "subCmdB", "subCmdB"},
		{"find subCmdB-1", "subCmdB subCmdB-1", "subCmdB-1"},
		{"find cmd404", "cmd404", ""},
		{"find subCmdC with parameter value a", "subCmdC arg1", "subCmdC"},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%v - %v", tc.name, tc.command)

		t.Run(testname, func(t *testing.T) {
			cmd := rootCommand.findCommand(strings.Split(tc.command, " "))

			if (cmd == nil && tc.expected != "") || (tc.expected != "" && cmd.Name != tc.expected) {
				t.Errorf("got: %v; wanted: %v", nil, tc.expected)
			}

			if (cmd != nil && tc.expected != "") && cmd.Name != tc.expected {
				t.Errorf("got: %v; wanted: %v", cmd.Name, tc.expected)
			}
		})
	}
}
