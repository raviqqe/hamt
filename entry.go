package hamt

// Entry represents an entry in a collection, which can be compared to values of
// type T (which is typically also the underlying type of the Entry).
type Entry[T any] interface {
	Hash() uint32
	Equal(T) bool
}
