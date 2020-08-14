package ast

const (
	// AST Node types
	Operator  = "op"
	Keyword   = "kw"
	Block     = "block"
	Constant  = "const"
	Reference = "ref"

	// error codes
	nilNodeErr      = "nil node passed onto stack operation"
	emptyStackErr   = "attempted to pop from empty stack"
	missingParenErr = "missing an opening parenthesis"

	// node type evaluation regexp patterns
	varNameRegexp = `^#?[a-zA-Z_]+$`
	numberRegexp  = `^[0-9]*\.?[0-9]*$`
)
