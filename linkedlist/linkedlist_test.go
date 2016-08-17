package linkedlist

import "testing"

func TestNode_Length(t *testing.T) {
	list := New()
	n1 := &Node{data: "n1"}
	n2 := &Node{data: "n2"}

	list.next = n1
	n1.next = n2

	if list.Length() != 2 {
		t.Errorf("wrong length, expected %d but got %d", 3, list.Length())
	}

	if d1, ok := n1.data.(string); !ok || d1 != "n1" {
		t.Errorf("node 1 wrong data, expected %s but got %s", "n1", d1)
	}

	if list.next != n1 {
		t.Error("head to node 1 connection is broken")
	}
	if d2, ok := n2.data.(string); !ok || d2 != "n2" {
		t.Errorf("node 2 wrong data, expected %s but got %s", "n2", d2)
	}
	if n1.next != n2 {
		t.Error("node 1 to node 2 connection is broken")
	}
}

func TestNode_Add(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.Length() != 3 {
		t.Errorf("wrong list length, expected %d but got %d", 3, list.Length())
	}

	if list.String() != "[1,2,3]" {
		t.Errorf("list has wrong order, expects %s but got %s", "[1,2,3]", list.String())
	}
}

func TestNode_Push(t *testing.T) {
	list := New()
	list.Push(1)
	list.Push(2)
	list.Push(3)

	if list.Length() != 3 {
		t.Errorf("wrong list length, expected %d but got %d", 3, list.Length())
	}

	if list.String() != "[3,2,1]" {
		t.Errorf("list has wrong order, expects %s but got %s", "[3,2,1]", list.String())
	}
}

func TestNode_Next(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if d, ok := list.Next().Next().Next().data.(int); !ok || d != 3 {
		t.Errorf("wrong data, expected %d but got %d", 3, d)
	}
}
