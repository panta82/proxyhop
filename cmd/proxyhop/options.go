package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
	. "proxyhop/tools"
)

type argumentsDefinition struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose log messages"`

	Port string `short:"p" long:"port" description:"Port to listen on" default:"10000"`

	Positional struct {
		Target string `description:"Target URL for proxying. URL paths will be appended" positional-arg-name:"TARGET"`
	} `positional-args:"yes" required:"1"`

	Help bool `short:"h" long:"help" description:"Show this help message"`
}

type Options struct {
	Verbose int
	Port string
	Target string
}

func loadOptions(args []string) Options {
	var parsedArgs argumentsDefinition

	parser := flags.NewParser(&parsedArgs, flags.PassDoubleDash)

	_, err := parser.ParseArgs(args)

	if parsedArgs.Help {
		parser.WriteHelp(os.Stdout)
		fmt.Printf(`
Examples:
  %s
  Proxy requests from http://localhost:12345/$x to https://google.com/$x  
`, EmText("> proxyhop -p 12345 https://google.com"))
		os.Exit(0)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, DangerText(err.Error()))
		os.Exit(1)
	}

	opts := Options{
		Verbose: len(parsedArgs.Verbose),
		Port: parsedArgs.Port,
		Target: parsedArgs.Positional.Target,
	}

	return opts
}

