package main

import (
	"fmt"
	"kumpmati/cscr"
)

func main() {
	compactScript := cscr.New()
	err := compactScript.Init()
	if err != nil {
		return
	}

	err = compactScript.Run()
	if err != nil {
		fmt.Printf("error running: %v\n", err)
		return
	}
}
