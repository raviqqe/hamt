package hamt

const arityBits = 5
const arity = 32

// Hamt represents a HAMT data structure.
type Hamt struct {
	level    uint8
	contents [arity]interface{}
}

// NewHamt creates a new HAMT.
func NewHamt(level uint8) Hamt {
	return Hamt{level: level}
}

// Insert inserts a value into a HAMT.
func (h Hamt) Insert(e Entry) Node {
	i := (e.Key() >> uint(arityBits*h.level)) % arity
	var c interface{}

	switch x := h.contents[i].(type) {
	case nil:
		c = e
	case Entry:
		l := h.level + 1

		if l*arityBits > arity {
			c = newBucket([]Entry{x, e})
		} else {
			c = NewHamt(l).Insert(e)
		}
	case Node:
		c = x.Insert(e)
	}

	g := h
	g.contents[i] = c
	return g
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
	// Traverse entries and sub nodes separately for cache locality.
	for _, c := range h.contents {
		if e, ok := c.(Entry); ok {
			return e, h.Delete(e)
		}
	}

	for i, c := range h.contents {
		if n, ok := c.(Node); ok {
			e, n := n.FirstRest()

			g := h
			g.contents[i] = n

			return e, g
		}
	}

	return nil, h // There is no entry inside. (h.Size() == 0)
}

// Size returns a size of a HAMT.
func (h Hamt) Size() int {
	s := 0

	for _, c := range h.contents {
		switch x := c.(type) {
		case Entry:
			s++
		case Node:
			s += x.Size()
		}
	}

	return s
}
