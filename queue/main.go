package queue

import "fmt"

// Queue struct
type Queue struct {
	Storage []Element
}

// Element struct
type Element struct {
	Value interface{}
}

// Enqueue item into queue
func (q *Queue) Enqueue(v interface{}) {
	e := Element{Value: v}
	q.Storage = append(q.Storage, e)
}

// Dequeue item from queue
func (q *Queue) Dequeue() interface{} {
	out, s := q.Storage[0], q.Storage[1:]
	q.Storage = s
	return out.Value
}

// Size returns the size of the queue's storage
func (q *Queue) Size() int {
	return len(q.Storage)
}

// Print each element in the queue
func (q *Queue) Print() {
	for _, e := range q.Storage {
		fmt.Println(e.Value)
	}
}
