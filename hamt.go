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
		if x.Equal(e) {
			return h.setChild(i, e)
		}

		l := h.level + 1

		if l*arityBits > arity {
			c = newBucket().Insert(x).Insert(e)
		} else {
			c = newHamt(l).Insert(x).Insert(e)
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

		switch n.State() {
		case empty:
			panic("Invariant error: trees must be normalized.")
		case singleton:
			e, _ := n.FirstRest()
			c = e
		}

		return h.setChild(i, c), true
	}

	return h, false
}

// Find finds a value in a HAMT.
func (h hamt) Find(e Entry) Entry {
	switch x := h.children[h.calculateIndex(e)].(type) {
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
			var e Entry
			e, n = n.FirstRest()

			if e != nil {
				return e, h.setChild(i, n)
			}
		}
	}

	return nil, h // There is no entry inside.
}

func (h hamt) ForEach(cb func(Entry) error) error {
	for _, child := range h.children {
		switch x := child.(type) {
		case nil:
			continue
		case Entry:
			if err := cb(x); err != nil {
				return err
			}
		case node:
			if err := x.ForEach(cb); err != nil {
				return err
			}
		}
	}
	return nil
}

// State returns a state of a HAMT.
func (h hamt) State() nodeState {
	es := 0
	ns := 0

	for _, c := range h.children {
		switch c.(type) {
		case Entry:
			es++
		case node:
			ns++
		}
	}

	if es+ns == 0 {
		return empty
	} else if es == 1 && ns == 0 {
		return singleton
	}

	return more
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
	return int((e.Hash() >> uint(arityBits*h.level)) % arity)
}

func (h hamt) setChild(i int, c interface{}) hamt {
	g := h
	g.children[i] = c
	return g
}
