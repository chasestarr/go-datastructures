package queue

import "testing"

func TestString(t *testing.T) {
	q := New()
	values := []string{"zero", "one", "two"}
	for _, e := range values {
		q.Enqueue(e)
	}

	for i := 0; i < 3; i++ {
		if dequeue := q.Dequeue(); dequeue != values[i] {
			t.Fatalf("queue string test failed actual: %s, expected: %s", dequeue, values[i])
		}
	}
}

func TestInt(t *testing.T) {
	q := New()
	values := []int{0, 1, 2}
	for _, e := range values {
		q.Enqueue(e)
	}

	for i := 0; i < 3; i++ {
		if dequeue := q.Dequeue(); dequeue != values[i] {
			t.Fatalf("queue int test failed actual: %d, expected: %d", dequeue, values[i])
		}
	}
}
