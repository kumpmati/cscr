package ast

import (
	"github.com/kumpmati/cscr/pkg/args"
	"github.com/kumpmati/cscr/pkg/lex"
)

// AST generator function type definition
type TraverserFunc func(t []lex.Token) *Stack

type Config struct {
	Traverser TraverserFunc
}

// returns the default AST generator config
func DefaultConfig(a args.Args) (cfg Config) {
	cfg.Traverser = DefaultTraverserFunc
	return
}
