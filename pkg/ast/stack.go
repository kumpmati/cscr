package ast

import (
	"sync"
)

// concurrent safe node stack
type Stack struct {
	items []Node
	lock  sync.RWMutex
}

// all the stacks that the ast traverser uses
type traverserStacks struct {
	operators Stack
	// nodes in this stack will be returned after traverser function returns
	out Stack
}

type OperationFunc func(stacks *traverserStacks, n *Node, next OperationFunc) error

// initializes a new empty stack
func (s *Stack) New() *Stack {
	s.items = []Node{}
	return s
}

// adds a node to the end of the stack
func (s *Stack) Push(n Node) {
	s.lock.Lock()
	s.items = append(s.items, n)
	s.lock.Unlock()
}

// removes a node from the end of the stack and returns it
func (s *Stack) Pop() *Node {
	if len(s.items) < 1 {
		return nil
	}
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item
}

// gets the last item of the stack without removing it from the stack
func (s *Stack) FakePop() *Node {
	if len(s.items) < 1 {
		return nil
	}
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.lock.Unlock()
	return &item
}

// gets the size of the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// initializes a new traverser stack
func newTraverserStack() *traverserStacks {
	stacks := traverserStacks{}

	// initialize stacks
	stacks.operators = *stacks.operators.New()
	stacks.out = *stacks.out.New()

	return &stacks
}
