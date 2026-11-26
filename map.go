package hamt

// Map represents a map (associative array).
type Map[K Entry[K], V any] struct {
	set Set[keyValue[K, V]]
}

// NewMap creates a new map.
func NewMap[K Entry[K], V any]() Map[K, V] {
	return Map[K, V]{NewSet[keyValue[K, V]]()}
}

// Insert inserts a key-value pair into a map.
func (m Map[K, V]) Insert(k K, v V) Map[K, V] {
	return Map[K, V]{m.set.Insert(newKeyValue(k, v))}
}

// Delete deletes a pair of a key and a value from a map.
func (m Map[K, V]) Delete(k K) Map[K, V] {
	var zero V
	return Map[K, V]{m.set.Delete(newKeyValue(k, zero))}
}

// Find finds a value corresponding to a given key from a map.
// If no value is found ok will be false
func (m Map[K, V]) Find(k K) (_ V, ok bool) {
	var zero V
	e, ok := m.set.find(newKeyValue(k, zero))
	return e.value, ok
}

// Include returns true if a key-value pair corresponding with a given key is
// included in a map, or false otherwise.
func (m Map[K, V]) Include(k K) bool {
	_, ok := m.Find(k)
	return ok
}

// FirstRest returns a key-value pair in a map and a rest of the map.
// This method is useful for iteration.
// If the map is empty, ok will be false.
func (m Map[K, V]) FirstRest() (_ K, _ V, _ Map[K, V], ok bool) {
	e, s, ok := m.set.FirstRest()
	m = Map[K, V]{s}
	return e.key, e.value, m, ok
}

func (m Map[K, V]) ForEach(cb func(K, V) error) error {
	return m.set.ForEach(func(kv keyValue[K, V]) error {
		return cb(kv.key, kv.value)
	})
}

// Merge merges 2 maps into one.
func (m Map[K, V]) Merge(n Map[K, V]) Map[K, V] {
	return Map[K, V]{m.set.Merge(n.set)}
}

// Size returns a size of a map.
func (m Map[K, V]) Size() int {
	return m.set.Size()
}
