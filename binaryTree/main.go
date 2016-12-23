package binaryTree

// BST struct
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// New BST
func New(v int) *BST {
	return &BST{Value: v}
}

// Add item to BST
func (b *BST) Add(v int) {
	if v < b.Value {
		if b.Left == nil {
			b.Left = New(v)
		} else {
			b.Left.Add(v)
		}
	} else {
		if b.Right == nil {
			b.Right = New(v)
		} else {
			b.Right.Add(v)
		}
	}
}

// Search for value in BST
func (b *BST) Search(v int) *BST {
	var output *BST

	if b.Value == v {
		output = b
	}

	if v < b.Value {
		if b.Left != nil {
			output = b.Left.Search(v)
		}
	} else {
		if b.Right != nil {
			output = b.Right.Search(v)
		}
	}

	return output
}
