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
	Find(T) *T
	FirstRest() (*T, node[T])
	ForEach(func(T) error) error
	State() nodeState
	Size() int // for debugging
}
