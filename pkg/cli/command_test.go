package cli

import (
	"fmt"
	"testing"
)

// TODO: Implement unit tests.
func TestFindCommandToExecute(t *testing.T) {
	var tests = []struct {
		name     string
		command  string
		expected string
	}{
		{"root cmd", "cli", "cli"},
		{"1 subcommand", "cli a", "cli a"},
	}

	for _, tc := range tests {
		tc := tc
		testname := fmt.Sprintf("%v - %v", tc.name, tc.command)
		t.Run(testname, func(t *testing.T) {
			// ans := FindCommandToExecute(tt.a, tt.b)
			// if ans != tt.want {
			t.Parallel()
			t.Errorf("got: %v; wanted: %v", nil, tc.expected)
			// }
		})
	}
}
