package tree

import "fmt"

// Tree struct
type Tree struct {
	Value    interface{}
	Children []Tree
}

// add
// contains
// traverse
// remove

// Add tree as child of target tree
func (t *Tree) Add(v interface{}) {
	t.Children = append(t.Children, Tree{Value: v})
}

// Contains returns true if value is contained in tree
func (t *Tree) Contains(v interface{}) bool {
	output := false

	// base case
	if t.Value == v {
		output = true
	}

	for _, c := range t.Children {
		output = c.Contains(v)
	}

	return output
}

// Traverse tree in order print values
func (t *Tree) Traverse() {
	fmt.Println(t.Value)
	for _, c := range t.Children {
		c.Traverse()
	}
}

// Find and return tree with value v
func (t *Tree) Find(v interface{}) *Tree {
	// TODO
}
