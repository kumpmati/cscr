package cscr

import (
	"github.com/kumpmati/cscr/pkg/args"
	"github.com/kumpmati/cscr/pkg/lex"
)

// default configuration for
func DefaultConfig() (c Config, err error) {
	// parse using command line arguments
	a, err := args.Parse(args.CommandLineArgs())
	if err != nil {
		return
	}
	c.Arguments = a

	// set lexer config to lexer default config
	c.LexerConfig = lex.DefaultConfig(a)
	return
}
