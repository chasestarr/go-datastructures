package redBlack

import "github.com/chasestarr/dataStructures/comparator"

// An RBTree contains value and color
// information of a red-black tree
type RBTree struct {
	value  interface{}
	color  string
	left   *RBTree
	right  *RBTree
	parent *RBTree
	less   comparator.Less
}

// New takes in a root interface and
// returns a red-black tree with that value
func New(value interface{}, comparator comparator.Less) *RBTree {
	return &RBTree{value: value, less: comparator, color: "black"}
}

// Insert takes in an interface, adds it to
// the red-black tree, and rebalances if needed
func (t *RBTree) Insert(value interface{}) {
	if t.less(value, t.value) {
		t.insertLeft(value)
	} else {
		t.insertRight(value)
	}
}

func (t *RBTree) insertLeft(value interface{}) {
	if t.left == nil {
		t.left = t.newChild(value)
		t.left.balance()
	} else {
		t.left.Insert(value)
	}
}

func (t *RBTree) insertRight(value interface{}) {
	if t.right == nil {
		t.right = t.newChild(value)
		t.right.balance()
	} else {
		t.right.Insert(value)
	}
}

func (t *RBTree) newChild(value interface{}) *RBTree {
	return &RBTree{
		value:  value,
		parent: t,
		color:  "red",
		less:   t.less,
	}
}

// balance
func (t *RBTree) balance() {

}
