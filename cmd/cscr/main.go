package main

import (
	"fmt"
	"github.com/kumpmati/cscr"
)

func main() {
	cfg, err := cscr.DefaultConfig()
	if err != nil {
		fmt.Printf("error while initializing default config: %v\n", err)
		return
	}
	// create new cscr
	c := cscr.New()
	err = c.Init(cfg)
	if err != nil {
		fmt.Printf("initialization error: %v\n", err)
		return
	}

	// run cscr
	err = c.Run()
	if err != nil {
		fmt.Printf("runtime error: %v\n", err)
		return
	}
}
