// transforms raw text to an array of token
package lex

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// lexer
type Lexer struct {
	config      Config
	currentFile *os.File
	tokens      []Token
}

// lexer config
type Config struct {
	LineParser LineParser
	FilePath   string
}

type LineParser func(s string) []Token

// returns a new lexer instance without configuring it
func New() Lexer { return Lexer{} }

// initializes the lexer with the given config
func (l *Lexer) Init(cfg Config) (err error) {
	// open file
	err = l.Open(cfg.FilePath)
	if err != nil {
		fmt.Printf("error while initializing lexer: %v\n", err)
		return
	}
	l.config = cfg
	return
}

// opens a file for reading. Note: does not
// close the file, Close() must be called later
func (l *Lexer) Open(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error while setting input file: %v\n", err)
		return
	}

	l.currentFile = f
	return
}

// runs the lexer, parsing the text into token
func (l *Lexer) Run() (err error) {
	if l.config.FilePath == "" {
		return errors.New("no file path specified")
	}

	l.tokens, err = l.LexFile(l.config.FilePath)
	if err != nil {
		return err
	}

	for _, v := range l.tokens {
		fmt.Printf("token: '%v', type: %v\n", v.Value, v.Type)
	}
	return
}

// opens a file for reading and transforms its contents to token,
// closing the file at the end
func (l *Lexer) LexFile(path string) (tokens []Token, err error) {
	err = l.Open(path)
	if err != nil {
		return
	}

	tokens, err = l.Lex()
	if err != nil {
		return
	}

	return tokens, l.currentFile.Close()
}

// reads the lexer struct's file line by line
func (l *Lexer) Lex() (tokens []Token, err error) {
	if l.config.LineParser == nil {
		panic("no line parser specified")
	}

	scanner := bufio.NewScanner(l.currentFile)
	for scanner.Scan() {
		txtLine := scanner.Text()
		tokens = append(tokens, l.config.LineParser(txtLine)...)
	}
	return
}
