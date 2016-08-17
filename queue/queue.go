package queue

type node struct {
	data interface{}
	next *node
	prev *node
}

// Queue is a collection designed for holding elements prior to processing.
type Queue struct {
	first *node
	last  *node
}

// New allocates new empty queue.
func New() *Queue {
	return &Queue{}
}

// Enqueue inserts the specified element into the queue.
func (q *Queue) Enqueue(d interface{}) {
	n := &node{
		data: d,
	}

	if q.first == nil {
		q.first = n
	}
	if q.last == nil {
		q.last = n
	} else {
		q.last.next = n
		n.prev = q.last
		q.last = n
	}
}

// Dequeue retrieves and removes the head of this queue.
// It returns false if queue is empty.
func (q *Queue) Dequeue() (interface{}, bool) {
	r := q.first
	if r == nil {
		return nil, false
	}
	if r.next != nil {
		q.first = r.next
		q.first.prev = nil
	}
	return r.data, true
}
