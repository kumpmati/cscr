package lex

import "testing"

func TestLexer_Init(t *testing.T) {
	validConfig := Config{
		FilePath:   "../../testdata/variable.cscr",
		LineParser: DefaultLineParser,
	}

	l := New()
	err := l.Init(validConfig)
	if err != nil {
		t.Error("lexer init failed with valid config")
	}

	invalidConfig := Config{
		FilePath:   "./non/existent/file.cscr",
		LineParser: nil,
	}

	// reset lexer
	l = New()
	err = l.Init(invalidConfig)
	if err == nil {
		t.Error("lexer init didn't fail with invalid config")
	}
}
