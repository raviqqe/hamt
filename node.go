package hamt

type nodeState int8

const (
	empty nodeState = iota
	singleton
	more
)

// node represents a node in a HAMT.
type node interface {
	Insert(Entry) node
	Delete(Entry) (node, bool)
	Find(Entry) Entry
	FirstRest() (Entry, node)
	State() nodeState
	Size() int // for debugging
}
