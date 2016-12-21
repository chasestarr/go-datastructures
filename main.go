package main

import "github.com/chasestarr/dataStructures/stack"
import "fmt"

func testStack() {
	s := stack.Stack{}

	s.Push("hello")
	s.Push("world")
	s.Print()
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
}

func main() {
	testStack()
}
