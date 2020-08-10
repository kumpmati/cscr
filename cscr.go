package cscr

import (
	"fmt"
	"github.com/kumpmati/cscr/pkg/args"
	"github.com/kumpmati/cscr/pkg/lex"
)

type Config struct {
	Arguments   args.Args
	LexerConfig lex.Config
}

type Cscr struct {
	config Config
	lexer  lex.Lexer
}

// returns a new cscr instance
func New() Cscr { return Cscr{} }

// initializes cscr with the given config
func (c *Cscr) Init(cfg Config) (err error) {
	// parse arguments first
	c.config = cfg

	// create a new lexer
	c.lexer = lex.New()
	// initialize the lexer with the given config
	err = c.lexer.Init(cfg.LexerConfig)
	if err != nil {
		fmt.Printf("error while initializing cscr: %v\n", err)
	}
	return
}

// runs the lexer
func (c *Cscr) Run() (err error) {
	return c.lexer.Run()
}

// default configuration
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
