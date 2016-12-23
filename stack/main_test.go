package stack

import "testing"

func TestString(t *testing.T) {
	s := New()
	values := []string{"zero", "one", "two"}
	for _, e := range values {
		s.Push(e)
	}

	for i := 2; i >= 0; i-- {
		if pop := s.Pop(); pop != values[i] {
			t.Fatalf("stack string test failed actual: %s, expected: %s", pop, values[i])
		}
	}
}

func TestInt(t *testing.T) {
	s := New()
	values := []int{0, 1, 2}
	for _, e := range values {
		s.Push(e)
	}

	for i := 2; i >= 0; i-- {
		if pop := s.Pop(); pop != values[i] {
			t.Fatalf("stack int test failed actual: %d, expected: %d", pop, values[i])
		}
	}
}
