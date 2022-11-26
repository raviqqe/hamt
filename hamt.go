package hamt

const arityBits = 5
const arity = 32

// hamt represents a HAMT data structure.
type hamt[T Entry[T]] struct {
	level    uint8
	children [arity]any
}

// newHamt creates a new HAMT.
func newHamt[T Entry[T]](level uint8) hamt[T] {
	return hamt[T]{level: level}
}

// Insert inserts a value into a HAMT.
func (h hamt[T]) Insert(e T) node[T] {
	i := h.calculateIndex(e)
	var c interface{}

	switch x := h.children[i].(type) {
	case nil:
		c = e
	case T:
		if x.Equal(e) {
			return h.setChild(i, e)
		}

		l := h.level + 1

		if l*arityBits > arity {
			c = newBucket[T]().Insert(x).Insert(e)
		} else {
			c = newHamt[T](l).Insert(x).Insert(e)
		}
	case node[T]:
		c = x.Insert(e)
	}

	return h.setChild(i, c)
}

// Delete deletes a value from a HAMT.
func (h hamt[T]) Delete(e T) (node[T], bool) {
	i := h.calculateIndex(e)

	switch x := h.children[i].(type) {
	case T:
		if x.Equal(e) {
			return h.setChild(i, nil), true
		}
	case node[T]:
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
			c = *e
		}

		return h.setChild(i, c), true
	}

	return h, false
}

// Find finds a value in a HAMT.
func (h hamt[T]) Find(e T) *T {
	switch x := h.children[h.calculateIndex(e)].(type) {
	case T:
		if x.Equal(e) {
			return &x
		}
	case node[T]:
		return x.Find(e)
	}

	return nil
}

// FirstRest returns a pointer to the first value and a HAMT without it.
// If h is empty, the pointer will be nil.
func (h hamt[T]) FirstRest() (*T, node[T]) {
	// Traverse entries and sub nodes separately for cache locality.
	for _, c := range h.children {
		if e, ok := c.(T); ok {
			h, _ := h.Delete(e)
			return &e, h
		}
	}

	for i, c := range h.children {
		if n, ok := c.(node[T]); ok {
			var e *T
			e, n = n.FirstRest()

			if e != nil {
				return e, h.setChild(i, n)
			}
		}
	}

	return nil, h // There is no entry inside.
}

func (h hamt[T]) ForEach(cb func(T) error) error {
	for _, child := range h.children {
		switch x := child.(type) {
		case nil:
			continue
		case T:
			if err := cb(x); err != nil {
				return err
			}
		case node[T]:
			if err := x.ForEach(cb); err != nil {
				return err
			}
		}
	}
	return nil
}

// State returns a state of a HAMT.
func (h hamt[T]) State() nodeState {
	es := 0
	ns := 0

	for _, c := range h.children {
		switch c.(type) {
		case T:
			es++
		case node[T]:
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
func (h hamt[T]) Size() int {
	s := 0

	for _, c := range h.children {
		switch x := c.(type) {
		case T:
			s++
		case node[T]:
			s += x.Size()
		}
	}

	return s
}

func (h hamt[T]) calculateIndex(e T) int {
	return int((e.Hash() >> uint(arityBits*h.level)) % arity)
}

func (h hamt[T]) setChild(i int, c any) hamt[T] {
	g := h
	g.children[i] = c
	return g
}
