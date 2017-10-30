package hamt

// Set represents a set.
type Set struct {
	size int
	hamt hamt
}

// NewSet creates a new set.
func NewSet() Set {
	return Set{0, newHamt(0)}
}

// Insert inserts a value into a set.
func (s Set) Insert(e Entry) Set {
	size := s.size

	if s.find(e) == nil {
		size++
	}

	return Set{size, s.hamt.Insert(e).(hamt)}
}

// Delete deletes a value from a set.
func (s Set) Delete(e Entry) Set {
	n, b := s.hamt.Delete(e)
	size := s.size

	if b {
		size--
	}

	return Set{size, n.(hamt)}
}

func (s Set) find(e Entry) Entry {
	return s.hamt.Find(e)
}

// Include returns true if a given entry is included in a set, or false otherwise.
func (s Set) Include(e Entry) bool {
	return s.find(e) != nil
}

// FirstRest returns a value in a set and a rest of the set.
// This method is useful for iteration.
func (s Set) FirstRest() (Entry, Set) {
	e, n := s.hamt.FirstRest()
	size := s.size

	if e != nil {
		size--
	}

	return e, Set{size, n.(hamt)}
}

// Size returns a size of a set.
func (s Set) Size() int {
	return s.size
}
