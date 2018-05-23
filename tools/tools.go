package tools

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func FatalError(message string, err *error) {
	fmt.Fprintf(os.Stderr,"%s %s\n", FatalText("FATAL"), DangerText(message))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s\n", FatalText("     "), DangerText((*err).Error()))

		stack := string(debug.Stack())
		for _, line := range strings.Split(stack, "\n") {
			fmt.Fprintf(os.Stderr, "%s %s\n", FatalText("     "), line)
		}
	}
	os.Exit(1)
}