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
			FilePath: "testdata/validCode.cscr",
		},
		LexerConfig: lex.Config{
			LineParser: nil,
			FilePath:   "testdata/validCode.cscr",
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
			LineParser: nil,
		},
	}

	// reset cscr
	c = New()
	err = c.Init(invalidConfig)
	if err == nil {
		t.Error("cscr didn't fail with invalid config")
	}
}
