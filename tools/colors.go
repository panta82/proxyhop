package tools

import "github.com/fatih/color"

var EmText func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()
var DangerText func(a ...interface{}) string = color.New(color.FgRed, color.Bold).SprintFunc()
var FatalText func(a ...interface{}) string = color.New(color.BgRed, color.FgWhite, color.Bold).SprintFunc()
