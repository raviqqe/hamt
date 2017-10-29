package hamt

// node represents a node in a HAMT.
type node interface {
	Insert(Entry) node
	Delete(Entry) (node, bool)
	Find(Entry) Entry
	FirstRest() (Entry, node)
	Size() int
}
