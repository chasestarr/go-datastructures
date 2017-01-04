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

// New takes in a root empty interface value and
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
	for t.parent != nil && t.parent.color == "red" {
		if t.isParentLeft() {
			t = t.balanceLeft()
		} else {
			t = t.balanceRight()
		}
	}

	t.root().color = "black"
}

func (t *RBTree) balanceLeft() *RBTree {
	uncle := t.rightUncle()
	if uncle != nil && uncle.color == "red" {
		uncle.color = "black"
		t.parent.color = "black"
		t.grandparent().color = "red"
		return t.grandparent()
	} else if !t.isLeft() {
		t = t.parent
		t.rotateLeft()
		t.parent.color = "black"
		t.grandparent().color = "red"
		t.grandparent().rotateRight()
	}
	return t
}

func (t *RBTree) balanceRight() *RBTree {
	uncle := t.leftUncle()
	if uncle != nil && uncle.color == "red" {
		uncle.color = "black"
		t.parent.color = "black"
		t.grandparent().color = "red"
		return t.grandparent()
	} else if t.isLeft() {
		t = t.parent
		t.rotateRight()
		t.parent.color = "black"
		t.grandparent().color = "red"
		t.grandparent().rotateLeft()
	}
	return t
}

func (t *RBTree) grandparent() *RBTree {
	return t.parent.parent
}

func (t *RBTree) leftUncle() *RBTree {
	return t.grandparent().left
}

func (t *RBTree) rightUncle() *RBTree {
	return t.grandparent().right
}

func (t *RBTree) rotateLeft() {
	y := t.right

	t.right = y.left
	if t.right != nil {
		t.left.parent = t
	}
	y.parent = t.parent
	if t.parent != nil && t.parent.left == t {
		t.parent.left = y
	} else if t.parent != nil && t.parent.right == t {
		t.parent.right = y
	}

	y.left = t
	t.parent = y
}

func (t *RBTree) rotateRight() {
	x := t.left

	t.left = x.right
	if t.left != nil {
		t.right.parent = t
	}
	x.parent = t.parent
	if t.parent != nil && t.parent.left == t {
		x.parent.left = x
	} else if t.parent != nil && t.parent.right == t {
		x.parent.right = x
	}
}

func (t *RBTree) isParentLeft() bool {
	if t.grandparent() != nil && t.parent == t.grandparent().left {
		return true
	}
	return false
}

func (t *RBTree) isLeft() bool {
	if t.parent != nil && t == t.parent.left {
		return true
	}
	return false
}

func (t *RBTree) root() *RBTree {
	current := t

	for current.parent != nil {
		current = current.parent
	}

	return current
}
