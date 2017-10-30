package hamt

// Entry represents an entry in a collection.
type Entry interface {
	Hash() uint32
	Equal(Entry) bool
}
