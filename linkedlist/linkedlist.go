package linkedlist

import (
	"bytes"
	"fmt"
)

// Node ...
type Node struct {
	data interface{}
	next *Node
}

// New allocates new head of the list.
func New() *Node {
	return &Node{}
}

// Length returns number of elements in the list.
func (n *Node) Length() (l int) {
	if n == nil {
		return
	}

	cn := n

	for {
		if cn.next == nil {
			break
		}

		l++
		cn = cn.next
	}

	return
}

// Push adds new node at the beginning of the list.
func (n *Node) Push(d interface{}) {
	nn := &Node{data: d, next: n.next}
	n.next = nn
}

// Add adds new node at the end of the list.
func (n *Node) Add(d interface{}) {
	cn := n

	for {
		if cn.next == nil {
			cn.next = &Node{
				data: d,
			}
			break
		}

		cn = cn.next
	}
}

// String implements fmt.Stringer interface.
func (n *Node) String() string {
	buf := bytes.NewBuffer(nil)
	cn := n

	for {
		if cn.next == nil {
			break
		}
		if buf.Len() != 0 {
			fmt.Fprint(buf, ",")
		} else {
			fmt.Fprint(buf, "[")

		}
		fmt.Fprintf(buf, "%v", cn.next.data)
		cn = cn.next
	}
	fmt.Fprint(buf, "]")

	return buf.String()
}

// Next return next node.
func (n *Node) Next() *Node {
	return n.next
}
