// transforms raw text to an array of token
package lex

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// lexer
type L struct {
	config      Config
	currentFile *os.File
	tokens      []Token
}

// returns a new lexer instance without initializing it
func New() L { return L{} }

// initializes the lexer with the given config
func (l *L) Init(cfg Config) (err error) {
	// open file
	err = l.Open(cfg.FilePath)
	if err != nil {
		return errors.New("file could not be opened")
	}

	if cfg.LineParser == nil {
		return errors.New("no line parser specified")
	}

	l.config = cfg
	return
}

// Opens a file for reading.
// NOTE: does not close the file, Close() must be called on the file later!
func (l *L) Open(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error while setting input file: %v\n", err)
		return
	}

	l.currentFile = f
	return
}

// runs the lexer, parsing the text into token
func (l *L) Run() (err error) {
	if l.config.FilePath == "" {
		return errors.New("no file path specified")
	}

	l.tokens, err = l.LexFile(l.config.FilePath)
	if err != nil {
		return err
	}
	return
}

func (l *L) GetTokens() []Token {
	return l.tokens
}

// Opens a file for reading and transforms its contents to tokens.
// Closes the opened file after completing.
func (l *L) LexFile(path string) (tokens []Token, err error) {
	err = l.Open(path)
	defer l.currentFile.Close()
	if err != nil {
		return
	}

	tokens, err = l.Lex()
	if err != nil {
		return
	}

	return
}

// reads the lexer struct's file line by line
func (l *L) Lex() (tokens []Token, err error) {
	scanner := bufio.NewScanner(l.currentFile)
	for scanner.Scan() {
		txtLine := scanner.Text()
		tokens = append(tokens, l.config.LineParser(txtLine)...)
	}
	return
}
