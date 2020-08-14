package cscr

import (
	"fmt"
	"github.com/kumpmati/cscr/pkg/args"
	"github.com/kumpmati/cscr/pkg/ast"
	"github.com/kumpmati/cscr/pkg/lex"
)

type Config struct {
	Arguments          args.Args
	LexerConfig        lex.Config
	AstGeneratorConfig ast.Config
}

type Cscr struct {
	config Config
	lexer  lex.Lexer
	ast    ast.Generator
}

// returns a new cscr instance
func New() Cscr { return Cscr{} }

// initializes cscr with the given config
func (c *Cscr) Init(cfg Config) (err error) {
	// parse arguments first
	c.config = cfg

	// lexer
	c.lexer = lex.New()
	err = c.lexer.Init(cfg.LexerConfig)
	if err != nil {
		return err
	}

	// ast generator
	c.ast = ast.New()
	err = c.ast.Init(cfg.AstGeneratorConfig)
	if err != nil {
		return err
	}
	return
}

// runs the lexer
func (c *Cscr) Run() (err error) {
	err = c.lexer.Run()
	if err != nil {
		return err
	}

	c.ast.SetTokens(c.lexer.GetTokens())
	err = c.ast.Run()
	if err != nil {
		return err
	}

	fmt.Println(*c.ast.GetTree())
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

	// set lexer config to default
	c.LexerConfig = lex.DefaultConfig(a)

	// set ast generator config to default
	c.AstGeneratorConfig = ast.DefaultConfig(a)
	return
}
