package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var options struct {
	Name    string `short:"n" default:"A name is more than Nomenclature"`
	Spanish bool   `short:"s" long:"spanish" description:"A name to say hello to."`
}

func main() {
	flags.Parse(&options)
	if options.Spanish == true {
		fmt.Printf("Names: %s", options.Name)
		return
	}
	fmt.Print("Ended")
}
