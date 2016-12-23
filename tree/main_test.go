package tree

import "testing"

func TestTree(t *testing.T) {
	tree := New(1)
	if tree.Contains(1) == false {
		t.Fatal("tree with one node should contain that value")
	}

	tree.Add(2)
	tree.Add(3)
	if tree.Contains(3) == false {
		t.Fatal("tree with one node should contain that value")
	}

	leaf := tree.Find(3)
	if leaf.Value != 3 {
		t.Fatal("Found tree with value should have target value")
	}

	leaf.Add(4)
	if tree.Contains(4) == false {
		t.Fatal("Should return true for leaf nodes")
	}
}
