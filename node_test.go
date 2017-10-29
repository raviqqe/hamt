package hamt

import "testing"

func TestHamtAsnode(t *testing.T) {
	t.Log(node(newHamt(0)))
}

func TestBucketAsnode(t *testing.T) {
	t.Log(node(newBucket(nil)))
}
