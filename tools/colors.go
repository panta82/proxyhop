package tools

import "github.com/fatih/color"

var DangerText func(a ...interface{}) string = color.New(color.FgRed, color.Bold).SprintFunc()

var EmText func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()