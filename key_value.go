package hamt

type keyValue[K Entry[K], V any] struct {
	key   K
	value V
}

func newKeyValue[K Entry[K], V any](k K, v V) keyValue[K, V] {
	return keyValue[K, V]{k, v}
}

func (kv keyValue[K, V]) Hash() uint32 {
	return kv.key.Hash()
}

func (kv keyValue[K, V]) Equal(e keyValue[K, V]) bool {
	return kv.key.Equal(e.key)
}
