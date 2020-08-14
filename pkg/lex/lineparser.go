package lex

import (
	"strings"
)

// default line parser
func DefaultLineParser(s string) (tokens []Token) {
	if len(s) == 0 {
		return
	}

	prevChar := string(s[0])
	currentLine := prevChar

	for _, v := range s[1:] {
		c := string(v)
		curr, prev := GetToken(c), GetToken(currentLine)
		if !prev.IsChainableWith(curr) {
			// create token only if current line contains non-whitespace characters
			if trimmedLine := strings.TrimSpace(currentLine); trimmedLine != "" {
				tokens = append(tokens, CreateToken(trimmedLine))
			}
			currentLine = ""
		}
		// append current line
		currentLine += c
		prevChar = c
	}

	if strings.TrimSpace(currentLine) != "" {
		tokens = append(tokens, CreateToken(currentLine))
	}
	// create newline token after everything else
	if len(tokens) > 0 && tokens[len(tokens)-1].Type != Break {
		tokens = append(tokens, CreateToken("\n"))
	}
	return
}
