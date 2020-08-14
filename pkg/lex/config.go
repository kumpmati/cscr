package lex

import "github.com/kumpmati/cscr/pkg/args"

// lexer config
type Config struct {
	LineParser LineParser
	FilePath   string
}

// returns the default lexer configuration
func DefaultConfig(a args.Args) (cfg Config) {
	// use default line parser
	cfg.LineParser = DefaultLineParser
	cfg.FilePath = a.FilePath
	return
}
