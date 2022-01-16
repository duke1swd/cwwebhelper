/*
 * Some useful code to access the CW website database
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	debug bool
)

func init() {
	flag.BoolVar(&debug, "D", false, "enable debugging")
	flag.Parse()
}

func usage() {
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	if flag.NArg() < 1 {
		usage()
	}

	switch flag.Arg(0) {
	case "users":
		usersCmd()
	default:
		fmt.Println("Unknown command")
		usage()
	}
}
