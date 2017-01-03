package redBlack

// Tree contains value and color information
// of a red-black tree
type Tree struct {
	value interface{}
	color string
}

// New takes in a root tree value and
// returns a red-black tree with that value
func New(value interface{}) *Tree {
	return &Tree{value: value, color: "black"}
}
