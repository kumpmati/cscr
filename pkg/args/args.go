// parses command line arguments
package args

import (
	"errors"
	"fmt"
	"os"
)

type Args struct {
	FilePath string
}

func Parse() (a Args, err error) {
	osArgs := os.Args[1:]
	if len(osArgs) == 0 {
		fmt.Println("type 'cscr help' to learn more about usage")
		return a, errors.New("no file specified in args")
	}

	if len(osArgs) == 1 {
		a.FilePath = osArgs[0]
		return
	}

	switch osArgs[0] {
	case "help":
		printHelp()
		break
	case "ver":
		printVersionNumber()
		break
	}

	return a, errors.New("not meant to run")
}

func printHelp() {
	fmt.Println("help is coming")
}

func printVersionNumber() {
	fmt.Println("version 0.0.1")
}
