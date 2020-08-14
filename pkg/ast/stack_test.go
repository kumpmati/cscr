package ast

import "testing"

var s Stack

func initStack() *Stack {
	if s.items == nil {
		s = Stack{}
		s.New()
	}
	return &s
}

func TestStack_Push(t *testing.T) {
	s := initStack()
	s.Push(Node{})
	s.Push(Node{})
	s.Push(Node{})

	if size := len(s.items); size != 3 {
		t.Errorf("wrong count, expected 3 but got %d", size)
	}
}

func TestStack_Pop(t *testing.T) {
	s := initStack()
	s.Push(Node{Value: "test"})

	if n := s.Pop(); n == nil || n.Value != "test" {
		t.Errorf("expected test node, got %v", n)
	}
}

// Fake popping should return the last element
// but should not remove it from the stack
func TestStack_FakePop(t *testing.T) {
	s := initStack()
	s.Push(Node{Value: "first"})
	s.Push(Node{Value: "test"})

	sizeBefore := s.Size()
	n := s.FakePop()
	sizeAfter := s.Size()

	// test that the size didn't change
	if sizeAfter != sizeBefore {
		t.Errorf("expected size %d, got %d", sizeBefore, sizeAfter)
	}

	// test that the correct item is returned
	if n == nil || n.Value != "test" {
		t.Errorf("expected test node, got %v", n)
	}
}
