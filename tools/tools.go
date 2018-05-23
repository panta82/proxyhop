package tools

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func doPrintError(label string, message string, err *error, stack string) {
	labelStandIn := strings.Repeat(" ", len(label))
	fmt.Fprintf(os.Stderr,"%s %s\n", DangerLabel(label), DangerText(message))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s\n", DangerLabel(labelStandIn), DangerText((*err).Error()))

		for index, line := range strings.Split(stack, "\n") {
			if index > 0 {
				fmt.Fprintf(os.Stderr, "%s %s\n", DangerLabel(labelStandIn), line)
			}
		}
	}
}

func PrintError(message string, err *error) {
	doPrintError("ERROR", message, err, string(debug.Stack()))
}

func FatalError(message string, err *error) {
	doPrintError("FATAL", message, err, string(debug.Stack()))
	os.Exit(1)
}