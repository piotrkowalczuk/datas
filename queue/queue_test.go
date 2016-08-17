package queue

import "testing"

func TestQueue(t *testing.T) {
	s := New()
	given := []int{1, 2, 3, 4, 5}

	for _, g := range given {
		s.Enqueue(g)
	}

	for i := range given {
		expected := given[i]
		v, ok := s.Dequeue()
		if !ok {
			t.Error("missing value in queue")
		}
		d, ok := v.(int)
		if !ok {
			t.Error("value in not an int")
		}
		if d != expected {
			t.Errorf("wrong value, expect %d but got %d", expected, d)
		}
	}
}
