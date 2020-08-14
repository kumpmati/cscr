package cscr

import (
	"github.com/kumpmati/cscr/pkg/args"
	"github.com/kumpmati/cscr/pkg/lex"
	"testing"
)

func TestCscr_Init(t *testing.T) {

	// test
	validConfig := Config{
		Arguments: args.Args{
			FilePath: "testdata/variable.cscr",
		},
		LexerConfig: lex.Config{
			LineParser: lex.DefaultLineParser,
			FilePath:   "testdata/variable.cscr",
		},
	}

	c := New()
	err := c.Init(validConfig)
	if err != nil {
		t.Error("cscr init failed with valid config")
	}

	invalidConfig := Config{
		Arguments: args.Args{
			FilePath: "",
		},
		LexerConfig: lex.Config{
			FilePath:   "",
			LineParser: nil,
		},
	}

	// reset cscr
	c = New()
	err = c.Init(invalidConfig)
	if err == nil {
		t.Error("expected cscr configuration to fail with invalid config")
	}
}
