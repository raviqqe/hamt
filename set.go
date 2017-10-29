package hamt

// Set represents a set.
type Set struct {
	hamt
}

// NewSet creates a new set.
func NewSet() Set {
	return Set{newHamt(0)}
}

// Insert inserts a value into a set.
func (s Set) Insert(e Entry) Set {
	return Set{s.hamt.Insert(e).(hamt)}
}

// Delete deletes a value from a set.
func (s Set) Delete(e Entry) Set {
	n, _ := s.hamt.Delete(e)
	return Set{n.(hamt)}
}

// Find finds a value in a set.
func (s Set) Find(e Entry) Entry {
	return s.hamt.Find(e)
}

// FirstRest returns a value in a set and a rest of the set.
// This method is useful for iteration.
func (s Set) FirstRest() (Entry, Set) {
	e, n := s.hamt.FirstRest()
	return e, Set{n.(hamt)}
}

// Size returns a size of a set.
func (s Set) Size() int {
	return s.hamt.Size()
}
