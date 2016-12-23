package main

import (
	"fmt"

	"github.com/chasestarr/dataStructures/hashTable"
	"github.com/chasestarr/dataStructures/queue"
	"github.com/chasestarr/dataStructures/stack"
)

func testStack() {
	s := stack.New()

	s.Push("hello")
	s.Push("world")
	s.Print()
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
}

func testQueue() {
	q := queue.New()

	q.Enqueue("Daffy")
	q.Enqueue("Duck")
	q.Print()
	fmt.Println(q.Dequeue())
	fmt.Println(q.Size())
	fmt.Println(q.Dequeue())
}

func testHash() {
	h := hashTable.New(20)
	h.Insert("dog", "corgi")
	h.Insert("dog", "basset hound")
	fmt.Println(h.Retrieve("dog"))
}

func main() {
	testStack()
	fmt.Println("================")
	testQueue()
	fmt.Println("================")
	testHash()
}
