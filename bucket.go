package hamt

type bucket[T Entry[T]] []T

func newBucket[T Entry[T]]() bucket[T] {
	return bucket[T](nil)
}

func (b bucket[T]) Insert(e T) node[T] {
	for i, f := range b {
		if e.Equal(f) {
			new := make(bucket[T], len(b))
			copy(new, b)
			new[i] = e
			return new
		}
	}

	return append(b, e)
}

func (b bucket[T]) Find(e T) (_ T, ok bool) {
	for _, f := range b {
		if e.Equal(f) {
			return f, true
		}
	}

	var ret T
	return ret, false
}

func (b bucket[T]) Delete(e T) (node[T], bool) {
	for i, f := range b {
		if e.Equal(f) {
			return append(b[:i], b[i+1:]...), true
		}
	}

	return b, false
}

func (b bucket[T]) FirstRest() (_ T, _ node[T], ok bool) {
	if len(b) == 0 {
		var ret T
		return ret, b, false
	}

	return b[0], b[1:], true
}

func (b bucket[T]) ForEach(cb func(T) error) error {
	for _, e := range b {
		if err := cb(e); err != nil {
			return err
		}
	}
	return nil
}

func (b bucket[T]) State() nodeState {
	switch len(b) {
	case 0:
		return empty
	case 1:
		return singleton
	}

	return more
}

func (b bucket[T]) Size() int {
	return len(b)
}
