package hamt

import "testing"

func TestHamtAsNode(t *testing.T) {
	t.Log(Node(NewHamt()))
}

func TestBucketAsNode(t *testing.T) {
	t.Log(Node(newBucket()))
}
