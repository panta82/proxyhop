package main

import (
	"os"
	. "proxyhop/proxy"
	"proxyhop/tools"
)

func main() {
	options := loadOptions(os.Args[1:])

	proxy := Proxy {
		Target: options.Positional.Target,
		Port: options.Port,
		Verbosity: len(options.Verbosity),
		CORSBusting: !options.NoCORSBusting,
	}

	err := proxy.Start()
	if err != nil {
		tools.FatalError("Proxy has crashed", &err)
	}
}