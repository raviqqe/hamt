package hamt

type bucket []Entry

func newBucket() bucket {
	return bucket(nil)
}

func (b bucket) Insert(e Entry) node {
	for i, f := range b {
		if e.Equal(f) {
			new := make(bucket, len(b))
			copy(new, b)
			new[i] = e
			return new
		}
	}

	return append(b, e)
}

func (b bucket) Find(e Entry) Entry {
	for _, f := range b {
		if e.Equal(f) {
			return f
		}
	}

	return nil
}

func (b bucket) Delete(e Entry) (node, bool) {
	for i, f := range b {
		if e.Equal(f) {
			return append(b[:i], b[i+1:]...), true
		}
	}

	return b, false
}

func (b bucket) FirstRest() (Entry, node) {
	if len(b) == 0 {
		return nil, b
	}

	return b[0], b[1:]
}

func (b bucket) ForEach(cb func(Entry) error) error {
	for _, e := range b {
		if err := cb(e); err != nil {
			return err
		}
	}
	return nil
}

func (b bucket) State() nodeState {
	switch len(b) {
	case 0:
		return empty
	case 1:
		return singleton
	}

	return more
}

func (b bucket) Size() int {
	return len(b)
}
