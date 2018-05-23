package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
	. "proxyhop/tools"
)

type options struct {
	Verbosity []bool `short:"v" long:"verbose" description:"Show verbose log messages"`

	Port string `short:"p" long:"port" description:"Port to listen on" default:"10000"`

	Positional struct {
		Target string `description:"Target URL for proxying. URL paths will be appended" positional-arg-name:"TARGET"`
	} `positional-args:"yes" required:"1"`

	Help bool `short:"h" long:"help" description:"Show this help message"`
}

func loadOptions(args []string) options {
	var opts options

	parser := flags.NewParser(&opts, flags.PassDoubleDash)

	_, err := parser.ParseArgs(args)

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		fmt.Printf(`
Examples:
  %s
  Proxy requests from http://localhost:12345/$x to https://google.com/$x  
`, EmText("> proxyhop -p 12345 https://google.com"))
		os.Exit(0)
	}

	if err != nil {
		PrintError(err.Error(), nil)
		os.Exit(1)
	}

	return opts
}

