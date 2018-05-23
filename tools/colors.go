package tools

import "github.com/fatih/color"

var MutedText func(a ...interface{}) string = color.New(color.FgWhite).SprintFunc()
var EmText func(a ...interface{}) string = color.New(color.FgHiYellow).SprintFunc()
var DangerText func(a ...interface{}) string = color.New(color.FgRed, color.Bold).SprintFunc()
var DangerLabel func(a ...interface{}) string = color.New(color.BgRed, color.FgWhite, color.Bold).SprintFunc()
