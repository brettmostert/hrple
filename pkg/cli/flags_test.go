package cli

import (
	"testing"

	"github.com/brettmostert/hrple/pkg/testHelper"
)

func TestStripFlags(t *testing.T) {
	args := []string{"subCmdDWithParamsAndFlags", "abc", "-verbose", "dev", "-party", "-name=moo", "--lastname", "xyz"}

	expectedValue := []string{"subCmdDWithParamsAndFlags", "abc", "dev", "xyz"}

	argsWithoutFlags := stripFlags(args)

	if !testHelper.Equal(argsWithoutFlags, expectedValue) {
		t.Errorf("verbose expected: '%v', got: '%v'", expectedValue, argsWithoutFlags)
	}
}

func TestSettingAndGettingFlags(t *testing.T) {
	cli := setupCliForTesting()

	args := []string{"subCmdDWithParamsAndFlags", "abc", "-verbose", "-party", "-name=moo", "--lastname", "abc"}

	cmd := cli.findCommand(args)
	if cmd == nil {
		t.Errorf("error: cmd is nil, args %v", args)
	}

	err := cmd.parseFlags(args)
	if err != nil {
		t.Errorf("error: unable to parse flags %v", err)
	}

	verboseFlagExpectedValue := true
	partyFlagExpectedValue := true
	nameFlagExpectedValue := "moo"

	verboseFlagValue, _ := cmd.Flags().GetBool("verbose")
	partyFlagValue, _ := cmd.Flags().GetBool("party")
	nameFlagValue := cmd.Flags().Value("name")

	if verboseFlagValue != verboseFlagExpectedValue {
		t.Errorf("verbose expected: '%v', got: '%v'", verboseFlagExpectedValue, verboseFlagValue)
	}

	if partyFlagValue != partyFlagExpectedValue {
		t.Errorf("party expected: '%v', got: '%v'", partyFlagExpectedValue, partyFlagValue)
	}

	if nameFlagValue != nameFlagExpectedValue {
		t.Errorf("name expected: '%v', got: '%v'", nameFlagExpectedValue, nameFlagValue)
	}
}
