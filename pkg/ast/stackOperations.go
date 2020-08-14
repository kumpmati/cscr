package ast

import (
	"errors"
)

// called every time an operator is pushed to the out stack
// should pop the last 2 elements from the out stack and add them as children to this node
func PushOperator(stacks *traverserStacks, n *Node) error {
	// need at least 2 elements in the out stack
	if stacks.out.Size() < 2 {
		return errors.New(emptyStackErr)
	}

	le, sle := popFromStack(&stacks.out), popFromStack(&stacks.out)
	if le != nil && sle != nil {
		n.Body = append(n.Body, *sle, *le)
		return PushToOutStack(stacks, n)
	}
	return nil
}

// pushes a node to a stack with basic error checking
func pushToStack(stack *Stack, n *Node) error {
	if n == nil {
		return errors.New(nilNodeErr)
	}
	stack.Push(*n)
	return nil
}

// pops a node from a stack with basic error checking
func popFromStack(stack *Stack) *Node {
	return stack.Pop()
}

// fake pops a node from a stack
func fakePopFromStack(stack *Stack) *Node {
	return stack.FakePop()
}

// pushes to operator stack and calls next function if defined
func PushToOperatorStack(stacks *traverserStacks, n *Node) error {
	return pushToStack(&stacks.operators, n)
}

// pushes to out stack and calls next function if defined
func PushToOutStack(stacks *traverserStacks, n *Node) error {
	return pushToStack(&stacks.out, n)
}

func popFromOperatorStack(stacks *traverserStacks) *Node {
	return popFromStack(&stacks.operators)
}

// returns true if b has equal or higher precedence than a
func hasHigherOrEqualPrecedence(a *Node, b *Node) bool {
	return b.Type.Precedence < a.Type.Precedence ||
		(b.Type.Precedence == a.Type.Precedence &&
			a.Type.Associate == "left")
}
