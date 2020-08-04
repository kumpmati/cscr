package cscr

import (
	"kumpmati/cscr/pkg/args"
	"kumpmati/cscr/pkg/lex"
)

type Cscr struct {
	arguments args.Args
	lexer     lex.Lexer
}

func New() Cscr {
	return Cscr{}
}

func (c *Cscr) Init() (err error) {
	// parse arguments
	a, err := args.Parse()
	if err != nil {
		return
	}

	c.arguments = a
	c.lexer = lex.New()
	return
}

func (c *Cscr) Run() (err error) {
	return c.lexer.Run(c.arguments)
}
