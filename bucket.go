package hamt

type bucket []Entry

func newBucket(es []Entry) bucket {
	return bucket(es)
}

func (b bucket) Insert(e Entry) Node {
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

func (b bucket) Delete(e Entry) (Node, bool) {
	for i, f := range b {
		if e.Equal(f) {
			return append(b[:i], b[i+1:]...), true
		}
	}

	return b, false
}

func (b bucket) FirstRest() (Entry, Node) {
	if b.Size() == 0 {
		return nil, b
	}

	return b[0], b[1:]
}

func (b bucket) Size() int {
	return len(b)
}
