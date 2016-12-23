package hashTable

import (
	"fmt"
	"hash/fnv"
)

// hash key
// insert
// remove
// retrieve
// log

// Pair --> key: value pair
type Pair struct {
	Key   string
	Value interface{}
}

// Bucket --> stores many pairs if collision
type Bucket struct {
	Pairs []Pair
}

// Hash --> top level storage for hash
type Hash struct {
	Buckets []Bucket
	Limit   int
}

func hash(s string, max int) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32()) % max
}

// New creates new hash object
func New(limit int) *Hash {
	b := make([]Bucket, limit)
	return &Hash{Buckets: b, Limit: limit}
}

// Insert key value pair to hash
func (h *Hash) Insert(key string, value string) {
	index := hash(key, h.Limit)
	pair := Pair{Key: key, Value: value}

	for _, p := range h.Buckets[index].Pairs {
		if p.Key == key {
			fmt.Println("same key")
			p.Value = value
			fmt.Println(p.Value)
			return
		}
	}

	h.Buckets[index].Pairs = append(h.Buckets[index].Pairs, pair)
}

// Retrieve item from hash
func (h *Hash) Retrieve(key string) interface{} {
	index := hash(key, h.Limit)
	for _, p := range h.Buckets[index].Pairs {
		if p.Key == key {
			return p.Value
		}
	}
	return nil
}

// Remove item from hash
func (h *Hash) Remove(key string) {
	index := hash(key, h.Limit)
	p := h.Buckets[index].Pairs
	var pair int
	for i, kv := range p {
		if kv.Key == key {
			pair = i
			break
		}
	}
	p[pair] = p[len(p)-1]
	p = p[:len(p)-1]
}
