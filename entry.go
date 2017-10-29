package champ

// Entry represents a hashable type.
type Entry interface {
	Key() int32
	Equal(Entry) bool
}
