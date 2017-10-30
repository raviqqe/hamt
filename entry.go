package hamt

// Entry represents an entry in a HAMT.
type Entry interface {
	Hash() uint32
	Equal(Entry) bool
}
