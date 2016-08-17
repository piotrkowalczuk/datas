package stack

type node struct {
	data interface{}
	next *node
}

// Stack represents a last-in-first-out (LIFO) stack of objects.
// The usual push and pop operations are provided.
type Stack struct {
	top *node
}

// New allocates an empty stack.
func New() *Stack {
	return &Stack{}
}

// Pop removes the object at the top of this stack and returns that object as the value of this function.
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	n := s.top
	s.top = s.top.next
	return n.data
}

// Push adds an item onto the top of this stack.
func (s *Stack) Push(d interface{}) {
	nn := &node{data: d, next: s.top}
	s.top = nn
}
