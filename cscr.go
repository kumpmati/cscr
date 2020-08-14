package cscr

import (
	"github.com/kumpmati/cscr/internal/runtime"
	"github.com/kumpmati/cscr/pkg/args"
	"github.com/kumpmati/cscr/pkg/ast"
	"github.com/kumpmati/cscr/pkg/lex"
)

type Config struct {
	Arguments          args.Args
	LexerConfig        lex.Config
	AstGeneratorConfig ast.Config
	RuntimeConfig      runtime.Config
}

type Cscr struct {
	config  Config
	lexer   lex.L
	ast     ast.A
	runtime runtime.R
}

// returns a new cscr instance
func New() Cscr { return Cscr{} }

// initializes cscr with the given config
func (c *Cscr) Init(cfg Config) error {
	// parse arguments first
	c.config = cfg

	c.lexer = lex.New()
	c.ast = ast.New()
	c.runtime = runtime.New()

	if err := c.lexer.Init(cfg.LexerConfig); err != nil {
		return err
	}
	if err := c.ast.Init(cfg.AstGeneratorConfig); err != nil {
		return err
	}
	if err := c.runtime.Init(cfg.RuntimeConfig); err != nil {
		return err
	}

	return nil
}

// runs the lexer
func (c *Cscr) Run() (err error) {
	// run lexer
	if err := c.lexer.Run(); err != nil {
		return err
	}

	// set ast tokens
	c.ast.SetTokens(c.lexer.GetTokens())
	if err := c.ast.Run(); err != nil {
		return err
	}

	// set runtime program
	c.runtime.SetProgram(c.ast.GetTree())
	if err := c.runtime.Run(); err != nil {
		return err
	}
	return
}

// default configuration
func DefaultConfig() (c Config, err error) {
	// parse using command line arguments
	a, err := args.Parse(args.CommandLineArgs())
	if err != nil {
		return
	}
	c.Arguments = a

	// get default lexer config
	c.LexerConfig = lex.DefaultConfig(a)

	// get default ast config
	c.AstGeneratorConfig = ast.DefaultConfig(a)

	// get default runtime config
	c.RuntimeConfig = runtime.DefaultRuntimeConfig(a)
	return
}
