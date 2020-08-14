package ast

import (
	"errors"
	"github.com/kumpmati/cscr/pkg/lex"
)

// ast generator struct
type A struct {
	config Config
	tokens []lex.Token
	tree   Node
}

// returns a new AST parser instance without initializing it
func New() A { return A{} }

// initializes the AST generator with the given config
func (a *A) Init(cfg Config) error {
	a.config = cfg
	return nil
}

// sets the generator's tokens
func (a *A) SetTokens(t []lex.Token) {
	a.tokens = t
}

// returns a pointer to the generator's tree
func (a *A) GetTree() *Node {
	return &a.tree
}

// runs the AST generator
func (a *A) Run() (err error) {
	if len(a.tokens) == 0 {
		return errors.New("no tokens to parse")
	}

	// create the abstract syntax tree and
	// init children as a node slice
	startNode := Node{
		Value: "root",
		Type:  NodeType{Type: Block},
		// traverse nodes to get body
		Body: a.config.Traverser(a.tokens).items,
	}
	a.tree = startNode
	return nil
}
