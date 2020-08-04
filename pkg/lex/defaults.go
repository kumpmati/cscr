package lex

import "github.com/kumpmati/cscr/pkg/args"

// returns the default lexer configuration
func DefaultConfig(a args.Args) (cfg Config) {
	// use default line parser
	cfg.LineParser = DefaultLineParser
	cfg.FilePath = a.FilePath
	return
}

// default line parser
func DefaultLineParser(s string) []Token {
	t := Token{
		Value: s,
		Properties: TokenProperties{
			IsKeyword:  false,
			IsOperator: false,
		},
	}
	return []Token{t}
}
