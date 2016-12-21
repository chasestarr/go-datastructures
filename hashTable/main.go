package hashTable

// hash key
// insert
// remove
// retrieve
// log

// Pair --> key: value pair
type Pair struct {
	Key   string
	value interface{}
}

// Bucket --> storas many pairs if collision
type Bucket struct {
	Pairs []Pair
}

// Hash --> top level storage for hash
type Hash struct {
	Buckets []Bucket
}
