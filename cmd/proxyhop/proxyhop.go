package main

import (
	"os"
	"fmt"
)

func main() {
	options := loadOptions(os.Args[1:])

	fmt.Println(options)
}