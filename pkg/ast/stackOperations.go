package ast

import (
	"errors"
)

// called every time an operator is pushed to the out stack
// should pop the last 2 elements from the out stack and add them as children to this node
func pushOperator(stacks *traverserStacks, n *Node) error {
	// throw error if out stack doesn't have at least 2 nodes
	if stacks.out.Size() < 2 {
		return errors.New(emptyStackErr)
	}

	le, sle := popFromStack(&stacks.out), popFromStack(&stacks.out)
	if le != nil && sle != nil {
		n.Body = append(n.Body, *sle, *le)
		return pushToOutStack(stacks, n)
	}
	return nil
}

// pushes a node to a stack with basic error checking
func pushToStack(stack *Stack, n *Node) error {
	if n == nil || stack == nil {
		return errors.New(nilNodeErr)
	}
	stack.Push(*n)
	return nil
}

// pops a node from a stack and returns it
func popFromStack(stack *Stack) *Node {
	return stack.Pop()
}

// returns the top node of a stack without removing it
func fakePopFromStack(stack *Stack) *Node {
	return stack.FakePop()
}

// pushes a node to the operator stack
func pushToOperatorStack(stacks *traverserStacks, n *Node) error {
	return pushToStack(&stacks.operators, n)
}

// pushes a node to the out stack
func pushToOutStack(stacks *traverserStacks, n *Node) error {
	return pushToStack(&stacks.out, n)
}

// pops a node from the operator stack
func popFromOperatorStack(stacks *traverserStacks) *Node {
	return popFromStack(&stacks.operators)
}

// returns true if b is of higher precedence (lower number = higher precedence)
// or if a and b have same precedence but a associates left-to-right
func hasHigherOrEqualPrecedence(a *Node, b *Node) bool {
	return b.Type.Precedence < a.Type.Precedence ||
		(b.Type.Precedence == a.Type.Precedence &&
			a.Type.Associate == "left")
}
