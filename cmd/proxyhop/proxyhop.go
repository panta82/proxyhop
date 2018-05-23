package main

import (
	"os"
	. "proxyhop/proxy"
	"proxyhop/tools"
)

func main() {
	options := loadOptions(os.Args[1:])

	proxy := Proxy {
		Target: options.Target,
	}

	err := proxy.SendRequest()
	if err != nil {
		tools.FatalError("Proxy request failed", &err)
	}
}