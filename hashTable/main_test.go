package hashTable

import "testing"

func TestHash(t *testing.T) {
	h := New(20)
	h.Insert("dog", "corgi")

	if h.Retrieve("dog") != "corgi" {
		t.Fatal("Hash Retieve error, key: 'dog' should return: 'corgi'")
	}

	h.Insert("dog", "basset hound")

	if h.Retrieve("dog") != "basset hound" {
		t.Fatal("Hash Insert overwrite error, key: 'dog' should return: 'basset hound'")
	}

	h.Remove("dog")

	if h.Retrieve("dog") == "basset hound" {
		t.Fatal("Hash Remove error, key: 'dog' should have exist in hash")
	}
}
