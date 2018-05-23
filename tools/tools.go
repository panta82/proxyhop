package tools

import (
	"fmt"
	"os"
)

func FatalError(message string, err *error) {
	fmt.Fprintf(os.Stderr, "FATAL: " + message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "       " + (*err).Error())
	}
	os.Exit(1)
}