package hamt

// Entry represents an entry in a HAMT.
type Entry interface {
	Hash() int32
	Equal(Entry) bool
}
