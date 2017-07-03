package main

import (
	"fmt"
	"os"

	"github.com/pgier/hawkbuild/cmd"
)

func main() {
	Main(os.Args)
}

func printHelp() {
	fmt.Println("usage: hawkbuild <command> [args]")
	fmt.Println("The available commands are: ")
	fmt.Printf(" license - %s\n", cmd.LicenseCmdDescription)
}

func printVersion() {
	fmt.Printf("%v version %v\n", os.Args[0], "0.1")
}

// Main program entry point
func Main(args []string) {
	if len(args) < 2 {
		printHelp()
		return
	}

	switch args[1] {
	case "-h", "--help", "help":
		printHelp()
	case "-v", "--version", "version":
		printVersion()
	case "license", "licenses":
		cmd.LicenseCmd(args[2:])
	default:
		fmt.Printf(fmt.Sprintf("%q %q is not valid command.\n", args[0], args[1]))
		printHelp()
		os.Exit(1)
	}
}
