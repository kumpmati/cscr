package lex

// token struct
type Token struct {
	Value     string
	Type      TokenType
	ChainWith Types
}
type Types []TokenType

// list of tokens with defined types and chaining behaviour
var tokens = map[string]Token{
	// string
	"\"": {Type: Str, ChainWith: Types{None}},
	"`":  {Type: Str, ChainWith: Types{None}},

	// separators
	",": {Type: Separator, ChainWith: Types{None}},
	".": {Type: Separator, ChainWith: Types{None}},

	// assignment
	":=": {Type: Operator, ChainWith: Types{None}},
	":":  {Type: Operator, ChainWith: Types{Operator}},
	"=":  {Type: Operator, ChainWith: Types{Math}},

	// math symbols
	"+": {Type: Math, ChainWith: Types{Math}},
	"-": {Type: Math, ChainWith: Types{Math}},
	"/": {Type: Math, ChainWith: Types{Math}},
	"*": {Type: Math, ChainWith: Types{Math}},

	// math operators
	"+=": {Type: MathOperator, ChainWith: Types{None}},
	"-=": {Type: MathOperator, ChainWith: Types{None}},
	"/=": {Type: MathOperator, ChainWith: Types{None}},
	"*=": {Type: MathOperator, ChainWith: Types{None}},

	// logic operators
	">":  {Type: Logic, ChainWith: Types{Math}},
	">=": {Type: Logic, ChainWith: Types{None}},
	"<":  {Type: Logic, ChainWith: Types{Math}},
	"<=": {Type: Logic, ChainWith: Types{None}},
	"==": {Type: Logic, ChainWith: Types{None}},
	"!":  {Type: Logic, ChainWith: Types{Math}},
	"!=": {Type: Logic, ChainWith: Types{None}},

	"(": {Type: Block, ChainWith: Types{None}},
	")": {Type: Block, ChainWith: Types{None}},

	";": {Type: Break, ChainWith: Types{None}},

	// keywords
	"fn":     {Type: Keyword, ChainWith: Types{None}},
	"return": {Type: Keyword, ChainWith: Types{None}},
	"if":     {Type: Keyword, ChainWith: Types{None}},
	"else":   {Type: Keyword, ChainWith: Types{None}},
	"true":   {Type: Keyword, ChainWith: Types{None}},
	"false":  {Type: Keyword, ChainWith: Types{None}},
}

var defaultToken = Token{Type: Default, ChainWith: Types{Self, Separator}}

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
	// get existing token if found, default token otherwise
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
		return t.Type == with.Type || contains(t.ChainWith, with.Type)
	}
	return contains(t.ChainWith, with.Type)
}

func contains(strArr []TokenType, str TokenType) bool {
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
