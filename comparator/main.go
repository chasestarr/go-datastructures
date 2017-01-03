package comparator

// Less type to compare two empty interfaces
// returns true if a is less than b
type Less func(a, b interface{}) bool

// Int accepts two empty interface values and
// returns true if a is less than b
// inputs will be type asserted to int
func Int(a, b interface{}) bool {
	return a.(int) < b.(int)
}

// String accepts two empty interface values and
// returns true if a is less than b
// inputs will be type asserted to string
func String(a, b interface{}) bool {
	return a.(string) < b.(string)
}
