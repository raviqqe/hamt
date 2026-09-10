package hamt

type nodeState int8

const (
	empty nodeState = iota
	singleton
	more
)

// node represents a node in a HAMT.
type node[T Entry[T]] interface {
	Insert(T) node[T]
	Delete(T) (node[T], bool)
	Find(T) (_ T, ok bool)
	FirstRest() (_ T, _ node[T], ok bool)
	ForEach(func(T) error) error
	State() nodeState
	Size() int // for debugging
}
