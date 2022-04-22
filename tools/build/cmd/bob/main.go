package main

import (
	"fmt"
	"os"

	"github.com/brettmostert/hrple/go/pkg/errors/exitError"
	"github.com/brettmostert/hrple/tools/build/internal/commands"
)

func main() {
	e := commands.NewExecuter()
	if _, err := e.Execute(); err != nil {
		fmt.Printf("failed executing command with error (%v)\n", err)
		os.Exit(int(exitError.Failure))
	}
}
