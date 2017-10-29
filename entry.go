package hamt

// Entry represents an entry in a HAMT.
type Entry interface {
	Key() int32
	Equal(Entry) bool
}
