package lex

type TokenType string

// token types
const (
	Math         TokenType = "mth"
	Logic        TokenType = "lgc"
	Operator     TokenType = "opr"
	MathOperator TokenType = "mop"
	Keyword      TokenType = "kwd"
	Separator    TokenType = "sep"
	Block        TokenType = "blk"
	Break        TokenType = "brk"
	Str          TokenType = "str"
	Default      TokenType = "dft"
	None         TokenType = "_none_"
	Self         TokenType = "_self_"
)

// line parser function type
type LineParser func(s string) []Token
