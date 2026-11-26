package hamt

// Set represents a set.
type Set[T Entry[T]] struct {
	size int
	hamt hamt[T]
}

// NewSet creates a new set.
func NewSet[T Entry[T]]() Set[T] {
	return Set[T]{0, newHamt[T](0)}
}

// Insert inserts a value into a set.
func (s Set[T]) Insert(e T) Set[T] {
	size := s.size

	if !s.Include(e) {
		size++
	}

	return Set[T]{size, s.hamt.Insert(e).(hamt[T])}
}

// Delete deletes a value from a set.
func (s Set[T]) Delete(e T) Set[T] {
	n, b := s.hamt.Delete(e)
	size := s.size

	if b {
		size--
	}

	return Set[T]{size, n.(hamt[T])}
}

func (s Set[T]) find(e T) (_ T, ok bool) {
	return s.hamt.Find(e)
}

// Include returns true if a given entry is included in a set, or false otherwise.
func (s Set[T]) Include(e T) bool {
	_, ok := s.find(e)
	return ok
}

// FirstRest returns a a value in a set and a rest of the set.
// This method is useful for iteration. If the set is empty, ok will
// be false.
func (s Set[T]) FirstRest() (_ T, _ Set[T], ok bool) {
	e, n, ok := s.hamt.FirstRest()
	size := s.size

	if ok {
		size--
	}

	return e, Set[T]{size, n.(hamt[T])}, ok
}

func (s Set[T]) ForEach(cb func(T) error) error {
	return s.hamt.ForEach(cb)
}

// Merge merges 2 sets into one.
func (s Set[T]) Merge(t Set[T]) Set[T] {
	for t.Size() != 0 {
		var (
			e  T
			ok bool
		)
		e, t, ok = t.FirstRest()
		if ok {
			s = s.Insert(e)
		}
	}

	return s
}

// Size returns a size of a set.
func (s Set[T]) Size() int {
	return s.size
}
