package ast

import (
	"fmt"
	"github.com/kumpmati/cscr/pkg/lex"
	"regexp"
)

// AST node type struct
type NodeType struct {
	Type       string
	SubType    lex.TokenType
	Associate  string
	Precedence int
}

// new node struct
type Node struct {
	Value string
	Type  NodeType
	Body  []Node
}

// Creates a new node and assigns it a type
// and stack operation based on the type
func nodeFromToken(t lex.Token) (Node, error) {
	nodeType := nodeTypeFromToken(t)
	n := Node{
		Value: t.Value,
		Type:  nodeType,
	}
	return n, nil
}

func nodeTypeFromToken(t lex.Token) (n NodeType) {
	switch t.Type {
	case lex.Operator:
		// operators and math operators
		n.Type = Operator
		n.SubType = t.Type
		n.Associate = "right"
		n.Precedence = 4
	case lex.Math, lex.MathOperator, lex.Logic:
		n.Type = Operator
		n.SubType = t.Type
		n.Associate = "left"
		switch t.Type {
		case lex.Logic:
			n.Precedence = 4
			n.Associate = "right"
		case lex.Math:
			if t.Value == "*" || t.Value == "/" {
				n.Precedence = 2
			} else {
				n.Precedence = 3
			}
		case lex.Operator, lex.MathOperator:
			n.Precedence = 5
			n.Associate = "right"
		}
	case lex.Block:
		n.Type = Block
		n.SubType = t.Type
		n.Precedence = 1
	case lex.Default:
		isVarName, err := regexp.MatchString(varNameRegexp, t.Value)
		isNumber, err := regexp.MatchString(numberRegexp, t.Value)
		switch {
		case err != nil:
			n.Type = "undefined"
			n.SubType = t.Type
		case isVarName:
			n.Type = Reference
		case isNumber:
			n.Type = Constant
		}
	}
	return
}

// prints the JSON representation of a node with recursive calls to its children
func (n Node) String() string {
	body := ""
	if len(n.Body) > 0 {
		body = ", \"body\": ["
		for _, c := range n.Body {
			body += c.String()
		}
		body += "]"
	}
	return fmt.Sprintf("{\"value\":\"%v\"%v},\n", n.Value, body)
}
