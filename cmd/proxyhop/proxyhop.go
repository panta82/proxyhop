package main

import (
	"os"
	. "proxyhop/proxy"
	"proxyhop/tools"
)

var version string

func main() {
	if version == "" {
		version = "0.0"
	}

	options := loadOptions(os.Args[1:], version)

	proxy := Proxy {
		Target: options.Positional.Target,
		Port: options.Port,
		Verbosity: options.getVerbosity(),
		CORSBusting: !options.NoCORSBusting,
	}

	err := proxy.Start()
	if err != nil {
		tools.FatalError("Proxy has crashed", &err)
	}
}