package hamt

// Hamt represents a HAMT data structure.
type Hamt struct{}

// NewHamt creates a new HAMT.
func NewHamt() Hamt {
	return Hamt{}
}

// Insert inserts a value into a HAMT.
func (h Hamt) Insert(e Entry) Node {
	return h
}

// Delete deletes a value from a HAMT.
func (h Hamt) Delete(e Entry) Node {
	return h
}

// Find finds a value in a HAMT.
func (h Hamt) Find(e Entry) Entry {
	return nil
}

// FirstRest returns a first value and a HAMT without it.
func (h Hamt) FirstRest() (Entry, Node) {
	return nil, h
}

// Size returns a size of a HAMT.
func (h Hamt) Size() int {
	return 0
}
