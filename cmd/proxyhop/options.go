package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
	. "proxyhop/tools"
)

type options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose log messages (including request bodies)"`

	Quiet bool `short:"q" long:"quiet" description:"Disable log output"`

	Port string `short:"p" long:"port" description:"Port to listen on" default:"10000"`

	Positional struct {
		Target string `description:"Target URL for proxying. URL paths will be appended" positional-arg-name:"TARGET"`
	} `positional-args:"yes" required:"1"`

	NoCORSBusting bool `long:"--disable-cors-busting" description:"Disable special CORS handling, just pass the requests normally"`

	Version bool `short:"V" long:"version" description:"Show program version and exit"`

	Help bool `short:"h" long:"help" description:"Show this help message"`
}

func (opts options) getVerbosity() int {
	if opts.Verbose { return 2 }
	if opts.Quiet { return 0 }
	return 1
}

func loadOptions(args []string, version string) options {
	var opts options

	parser := flags.NewParser(&opts, flags.PassDoubleDash)

	_, err := parser.ParseArgs(args)

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		fmt.Printf(`
Example:
  %s
  (Proxy requests from http://localhost:12345/$x to https://google.com/$x)  
`, EmText("proxyhop -p 12345 https://google.com"))
		os.Exit(0)
	}

	if opts.Version {
		fmt.Printf("proxyhop %s\nby Pantas, 2018\n", version)
		os.Exit(0)
	}

	if err != nil {
		PrintError(err.Error(), nil)
		os.Exit(1)
	}

	return opts
}

