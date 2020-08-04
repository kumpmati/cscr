// parses command line arguments
package args

import (
	"errors"
	"fmt"
	"os"
)

// arguments structure
type Args struct {
	FilePath string
}

// return command line arguments
func CommandLineArgs() []string {
	return os.Args[1:]
}

// parses an array of strings into a args struct
func Parse(argArr []string) (a Args, err error) {
	if len(argArr) == 0 {
		fmt.Println("type 'cscr help' to learn more about usage")
		return a, errors.New("no file specified in args")
	}

	if len(argArr) == 1 {
		a.FilePath = argArr[0]
		return
	}

	switch argArr[0] {
	case "help":
		printHelp()
		break
	case "ver":
		printVersionNumber()
		break
	}

	return a, errors.New("not meant to run")
}

// prints the help message for cscr
func printHelp() {
	fmt.Println("help is coming")
}

// prints cscr's current version number
func printVersionNumber() {
	fmt.Println("version 0.0.1")
}
