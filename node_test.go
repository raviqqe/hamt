package hamt

import "testing"

func TestHamtAsNode(t *testing.T) {
	t.Log(node[entryInt](newHamt[entryInt](0)))
}

func TestBucketAsNode(t *testing.T) {
	t.Log(node[entryInt](newBucket[entryInt]()))
}
