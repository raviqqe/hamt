package hamt

import "testing"

func TestHamtAsNode(t *testing.T) {
	t.Log(Node(NewHamt(0)))
}

func TestBucketAsNode(t *testing.T) {
	t.Log(Node(newBucket(nil)))
}
