package hamt

const arityBits = 5
const arity = 32

// hamt represents a HAMT data structure.
type hamt struct {
	level    uint8
	children [arity]interface{}
}

// newHamt creates a new HAMT.
func newHamt(level uint8) hamt {
	return hamt{level: level}
}

// Insert inserts a value into a HAMT.
func (h hamt) Insert(e Entry) node {
	i := h.calculateIndex(e)
	var c interface{}

	switch x := h.children[i].(type) {
	case nil:
		c = e
	case Entry:
		l := h.level + 1

		if l*arityBits > arity {
			c = newBucket([]Entry{x, e})
		} else {
			c = newHamt(l).Insert(e)
		}
	case node:
		c = x.Insert(e)
	}

	return h.setChild(i, c)
}

// Delete deletes a value from a HAMT.
func (h hamt) Delete(e Entry) (node, bool) {
	i := h.calculateIndex(e)

	switch x := h.children[i].(type) {
	case Entry:
		if x.Equal(e) {
			return h.setChild(i, nil), true
		}
	case node:
		n, b := x.Delete(e)

		if !b {
			return h, false
		}

		var c interface{} = n

		switch n.Size() {
		case 0:
			c = nil
		case 1:
			e, _ := n.FirstRest()
			c = e
		}

		return h.setChild(i, c), true
	}

	return h, false
}

// Find finds a value in a HAMT.
func (h hamt) Find(e Entry) Entry {
	i := h.calculateIndex(e)

	switch x := h.children[i].(type) {
	case Entry:
		if x.Equal(e) {
			return x
		}
	case node:
		return x.Find(e)
	}

	return nil
}

// FirstRest returns a first value and a HAMT without it.
func (h hamt) FirstRest() (Entry, node) {
	// Traverse entries and sub nodes separately for cache locality.
	for _, c := range h.children {
		if e, ok := c.(Entry); ok {
			h, _ := h.Delete(e)
			return e, h
		}
	}

	for i, c := range h.children {
		if n, ok := c.(node); ok {
			e, n := n.FirstRest()
			return e, h.setChild(i, n)
		}
	}

	return nil, h // There is no entry inside. (h.Size() == 0)
}

// Size returns a size of a HAMT.
func (h hamt) Size() int {
	s := 0

	for _, c := range h.children {
		switch x := c.(type) {
		case Entry:
			s++
		case node:
			s += x.Size()
		}
	}

	return s
}

func (h hamt) calculateIndex(e Entry) int {
	return int((e.Key() >> uint(arityBits*h.level)) % arity)
}

func (h hamt) setChild(i int, c interface{}) hamt {
	g := h
	g.children[i] = c
	return g
}
