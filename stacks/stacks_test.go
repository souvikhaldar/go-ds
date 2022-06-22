package stacks

import (
	"testing"

	"gotest.tools/assert"
)

func TestPush(t *testing.T) {
	s := NewStacks()
	s.Push(3)
	s.Push(2)
	s.Push(1)
	tos, err := s.Top()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, tos, 1)
	el, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, el, 1)
}

//
