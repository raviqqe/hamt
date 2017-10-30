package hamt

type keyValue struct {
	key   Entry
	value interface{}
}

func newKeyValue(k Entry, v interface{}) keyValue {
	return keyValue{k, v}
}

func (kv keyValue) Hash() uint32 {
	return kv.key.Hash()
}

func (kv keyValue) Equal(e Entry) bool {
	if kv, ok := e.(keyValue); ok {
		e = kv.key
	}

	return kv.key.Equal(e)
}
