package cli

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brettmostert/hrple/pkg/testHelper"
)

func ExecuteTestCommand(cmd *Command, args []string) ([]interface{}, error) {
	returnValue := make([]interface{}, len(args))
	for i := range args {
		returnValue[i] = args[i]
	}

	return returnValue, nil
}

func SetupCliForTesting() *Command {
	rootCommand := &Command{
		Name: "cli",
		Run:  ExecuteTestCommand,
	}

	rootCommand.AddCommand(&Command{
		Name: "subCmdA",
		Run:  ExecuteTestCommand,
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
		Run:  ExecuteTestCommand,
	}

	rootCommand.AddCommand(subCmdC)

	return rootCommand
}

func TestFindCommand(t *testing.T) {
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
				t.Errorf("expected: '%v', got: '%v'", tc.expected, nil)
			}

			if (cmd != nil && tc.expected != "") && cmd.Name != tc.expected {
				t.Errorf("expected: '%v', got: '%v'", tc.expected, cmd.Name)
			}
		})
	}
}

func TestAddCommandCannotBeChildOfSelf(t *testing.T) {
	cmd := &Command{
		Name: "subCmdB",
	}

	add := func() {
		cmd.AddCommand(cmd)
	}

	testHelper.ShouldPanic(t, add, "Command cannot be a child of itself")
}

func TestAddCommandParentAndChildCannotHaveSameName(t *testing.T) {
	parent := &Command{
		Name: "parent",
	}

	add := func() {
		parent.AddCommand(&Command{
			Name: "parent",
		})
	}

	testHelper.ShouldPanic(t, add, "Parent command and child command cannot have the same name")
}

func TestExecuteParent(t *testing.T) {
	cli := SetupCliForTesting()

	var expected []interface{}

	returnValues, err := cli.Execute(Options{})

	if err != nil && len(returnValues) != len(expected) {
		t.Errorf("expected: '%v' with no errors, got: '%v', error: '%v'", expected, returnValues, err)
	}
}

func TestExecuteSubCommand(t *testing.T) {
	cli := SetupCliForTesting()

	returnValues, err := cli.Execute(Options{
		Args: []string{"subCmdA"},
	})

	var expected []interface{}

	if err != nil && len(returnValues) != len(expected) {
		t.Errorf("expected: '%v' with no errors, got: '%v', error: '%v'", expected, returnValues, err)
	}
}

func TestExecuteSubCommandWithParam(t *testing.T) {
	cli := SetupCliForTesting()

	returnValues, err := cli.Execute(Options{
		Args: []string{"subCmdC", "abc", "xyz"},
	})

	expected := make([]interface{}, 2)
	expected[0] = "abc"
	expected[1] = "xyz"

	if err != nil && len(returnValues) != len(expected) {
		t.Errorf("expected: '%v' with no errors, got: '%v', error: '%v'", expected, returnValues, err)
	}

	for i := range expected {
		if returnValues[i] != expected[i] {
			t.Errorf("expected: '%v' at index '%v', got: '%v'", returnValues[i], i, expected[i])
		}
	}
}

func TestExecuteCannotFindCommand(t *testing.T) {
	expected := "Command not found, args: subCmd404"
	cli := SetupCliForTesting()

	_, err := cli.Execute(Options{
		Args: []string{"subCmd404"},
	})

	if err == nil || err.Error() != expected {
		t.Errorf("got: '%v'; wanted the following error '%v'", err, expected)
	}
}
