package champ

type bucket []Entry

func newBucket() bucket {
	return nil
}

func (b bucket) Insert(k Entry) bucket {
	return append(b, k)
}

func (b bucket) Find(k Entry) Entry {
	for _, k := range b {
		if k.Equal(k) {
			return k
		}
	}

	return nil
}

func (b bucket) Delete(k Entry) bucket {
	for i, k := range b {
		if k.Equal(k) {
			return append(b[:i], b[i:]...)
		}
	}

	return b
}
