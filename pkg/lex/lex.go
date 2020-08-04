// transforms raw text to an array of tokens
package lex

import (
	"bufio"
	"errors"
	"kumpmati/cscr/pkg/args"
	"os"
)

// token struct
type Token struct {
	Value      string
	Properties TokenProperties
}

type TokenProperties struct {
	IsKeyword  bool
	IsOperator bool
}

// lexer
type Lexer struct {
	File       *os.File
	Tokens     []Token
	arguments  args.Args
	lineParser lineParser
}

type lineParser func(s string) []Token

func New() (l Lexer) {
	l = Lexer{}
	l.Defaults()
	return
}

func (l *Lexer) Defaults() {
	l.lineParser = parseLine
}

func (l *Lexer) Run(a args.Args) (err error) {
	if a.FilePath == "" {
		return errors.New("no file path specified")
	}
	l.arguments = a

	tokens, err := l.ReadFile(l.arguments.FilePath, l.lineParser)
	if err != nil {
		return err
	}

	l.Tokens = tokens
	return
}

// opens a file for reading
func (l *Lexer) SetInputFile(path string) (f *os.File, err error) {
	f, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	l.File = f
	return
}

func (l *Lexer) Read(lp lineParser) (tokens []Token, err error) {
	tokens = []Token{}
	if lp == nil {
		return tokens, errors.New("no lineparser specified")
	}

	scanner := bufio.NewScanner(l.File)
	for scanner.Scan() {
		tokens = append(tokens, lp(scanner.Text())...)
	}
	return
}

func (l *Lexer) ReadFile(path string, lp lineParser) (tokens []Token, err error) {
	f, err := l.SetInputFile(path)
	if err != nil {
		return
	}
	defer f.Close()

	tokens, err = l.Read(lp)
	if err != nil {
		return
	}

	return
}

func parseLine(s string) []Token {
	t := Token{
		Value: s,
		Properties: TokenProperties{
			IsKeyword:  false,
			IsOperator: false,
		},
	}
	return []Token{t}
}
