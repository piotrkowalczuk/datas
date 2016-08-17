package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	given := []int{1, 2, 3, 4, 5}

	for _, g := range given {
		s.Push(g)
	}

	for i := range given {
		expected := given[len(given)-i-1]
		if v, ok := s.Pop().(int); !ok || v != expected {
			t.Errorf("wrong first value, expect %d but got %d", expected, v)
		}
	}
}
