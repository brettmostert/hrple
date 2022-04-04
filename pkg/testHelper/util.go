package testHelper

import (
	"fmt"
	"testing"
)

func ShouldPanic(t *testing.T, f func(), panicMessage ...string) {
	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			recoveryMessageString := fmt.Sprint(recoveryMessage)

			if len(panicMessage) != 0 {
				msg := panicMessage[0]
				if msg != recoveryMessageString {
					t.Errorf("expected: '%v', got: '%v'", msg, recoveryMessage)
				}
			}
		}
	}()
	f()
	t.Errorf("should have panicked")
}
