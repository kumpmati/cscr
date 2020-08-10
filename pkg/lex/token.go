package lex

type Types []string

type Token struct {
	Value     string
	Type      string
	ChainWith Types
}

// token types
const (
	Operator     = "opn"
	MathOperator = "mop"
	Keyword      = "kwd"
	Separator    = "sep"
	Block        = "blk"
	Expression   = "exp"
	Accessor     = "acc"
	Break        = "brk"
	Str          = "str"
	Default      = "dft"
	None         = "_none_"
	Self         = "_self_"
)

// store all properties with their respective properties here
var tokens = map[string]Token{
	// string
	"\"": {Type: Str, ChainWith: Types{None}},
	"`":  {Type: Str, ChainWith: Types{None}},

	// assignment
	":=": {Type: Operator, ChainWith: Types{None}},
	":":  {Type: Operator, ChainWith: Types{MathOperator}},

	// separators
	",": {Type: Separator, ChainWith: Types{None}},
	".": {Type: Separator, ChainWith: Types{None}},

	// math operators
	"=":  {Type: MathOperator, ChainWith: Types{MathOperator}},
	"+":  {Type: MathOperator, ChainWith: Types{MathOperator}},
	"++": {Type: MathOperator, ChainWith: Types{None}},
	"-":  {Type: MathOperator, ChainWith: Types{MathOperator, Operator}},
	"--": {Type: MathOperator, ChainWith: Types{None}},
	"/":  {Type: MathOperator, ChainWith: Types{MathOperator}},
	"*":  {Type: MathOperator, ChainWith: Types{MathOperator}},

	// reassignment operators
	"+=": {Type: MathOperator, ChainWith: Types{None}},
	"-=": {Type: MathOperator, ChainWith: Types{None}},
	"*=": {Type: MathOperator, ChainWith: Types{None}},
	"/=": {Type: MathOperator, ChainWith: Types{None}},

	"(": {Type: Expression, ChainWith: Types{None}},
	")": {Type: Expression, ChainWith: Types{None}},

	"[": {Type: Accessor, ChainWith: Types{None}},
	"]": {Type: Accessor, ChainWith: Types{None}},

	"{": {Type: Block, ChainWith: Types{None}},
	"}": {Type: Block, ChainWith: Types{None}},

	// breaks
	" ":  {Type: Break, ChainWith: Types{None}},
	"\n": {Type: Break, ChainWith: Types{None}},
	";":  {Type: Break, ChainWith: Types{None}},

	// keywords
	"func":   {Type: Keyword, ChainWith: Types{None}},
	"return": {Type: Keyword, ChainWith: Types{None}},
	"if":     {Type: Keyword, ChainWith: Types{None}},
	"else":   {Type: Keyword, ChainWith: Types{None}},
}

var defaultToken = Token{Type: Default, ChainWith: Types{Self}}

// returns a token based on the given string
func GetToken(s string) (t Token) {
	t, exists := tokens[s]
	if !exists {
		t = defaultToken
	}
	return
}

// creates a new token. If a token is defined in the tokens map with the value of s,
// then the token will get the properties of that token. If a defined token isn't found,
// the new token gets the default properties
func CreateToken(s string) (t Token) {
	t = GetToken(s)
	t.Value = s
	return
}

// returns true if 'with' is chainable with 't'
func (t Token) IsChainableWith(with Token) bool {
	if contains(t.ChainWith, None) {
		return false
	}
	if contains(t.ChainWith, Self) {
		return t.Type == with.Type
	}
	return contains(t.ChainWith, with.Type)
}

func contains(strArr []string, str string) bool {
	if len(strArr) == 0 {
		return false
	}

	for _, v := range strArr {
		if v == str {
			return true
		}
	}
	return false
}
