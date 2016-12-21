package queue

// Queue struct
type Queue struct {
	Storage []string
}

// enqueue
// dequeue
// size
// print

// Enqueue item into queue
func (q *Queue) Enqueue(str string) {
	q.Storage = append(q.Storage, str)
}
