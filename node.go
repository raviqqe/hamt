package hamt

// Node represents a node in a HAMT.
type Node interface {
	Insert(Entry) Node
	Delete(Entry) (Node, bool)
	Find(Entry) Entry
	FirstRest() (Entry, Node)
	Size() int
}
