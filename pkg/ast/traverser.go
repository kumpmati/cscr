package ast

import (
	"github.com/kumpmati/cscr/pkg/lex"
)

// default token traverser function. turns tokens into an abstract syntax tree
func DefaultTraverserFunc(tokens []lex.Token) (out *Stack) {
	stacks := newTraverserStack()

	for _, t := range tokens {
		newNode, err := nodeFromToken(t)
		if err != nil {
			// skip to next node silently if node creation failed
			continue
		}
		switch newNode.Type.Type {
		case Constant, Reference:
			err = PushToOutStack(stacks, &newNode)
		case Operator:
			switch newNode.Type.SubType {
			default:
				topOperator := fakePopFromStack(&stacks.operators)
				// while last of operator stack is an operator,
				// not a opening parenthesis, and has higher or equal precedence (should be calculated first)
				for topOperator != nil && topOperator.Type.Type == Operator &&
					hasHigherOrEqualPrecedence(&newNode, topOperator) &&
					topOperator.Value != "(" {
					// pop to out stack from operator stack
					err = PushOperator(stacks, popFromOperatorStack(stacks))
					topOperator = fakePopFromStack(&stacks.operators)
				}
				// push top operator stack
				err = PushToOperatorStack(stacks, &newNode)
			}
		case Block:
			switch newNode.Value {
			case "(":
				err = PushToOperatorStack(stacks, &newNode)
			case ")":
				topOperator := fakePopFromStack(&stacks.operators)

				for topOperator != nil && topOperator.Value != "(" {
					err = PushOperator(stacks, popFromOperatorStack(stacks))
					topOperator = fakePopFromStack(&stacks.operators)

					if stacks.operators.Size() == 0 {
						panic(missingParenErr)
					}
				}

				if topOperator != nil && topOperator.Value == "(" {
					popFromOperatorStack(stacks)
				}
			}
		default:
			for stacks.operators.Size() > 0 {
				if node := popFromOperatorStack(stacks); node != nil {
					err = PushOperator(stacks, node)
				}
			}
		}

		if err != nil {
			panic(err)
		}
	}
	for stacks.operators.Size() > 0 {
		if node := popFromOperatorStack(stacks); node != nil {
			if err := PushOperator(stacks, node); err != nil {
				panic(err)
			}
		}
	}

	return &stacks.out
}
