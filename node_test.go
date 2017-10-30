package hamt

import "testing"

func TestHamtAsNode(t *testing.T) {
	t.Log(node(newHamt(0)))
}

func TestBucketAsNode(t *testing.T) {
	t.Log(node(newBucket()))
}
