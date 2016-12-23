package binaryTree

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	bst := New(5)
	bst.Add(2)
	bst.Add(7)

	two := bst.Search(2)
	fmt.Println(two == nil)
	if two.Value != 2 {
		t.Fatal("Should find value in BST")
	}

	seven := bst.Search(7)
	if seven.Value != 7 {
		t.Fatal("Should find value in BST")
	}
}
