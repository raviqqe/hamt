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
// It returns nil if no value is found.
func (m Map[K, V]) Find(k K) *V {
	var zero V
	e := m.set.find(newKeyValue(k, zero))

	if e == nil {
		return nil
	}

	return &e.value
}

// Include returns true if a key-value pair corresponding with a given key is
// included in a map, or false otherwise.
func (m Map[K, V]) Include(k K) bool {
	return m.Find(k) != nil
}

// FirstRest returns a key-value pair in a map and a rest of the map.
// This method is useful for iteration.
// The key and value would be nil if the map is empty.
func (m Map[K, V]) FirstRest() (*K, *V, Map[K, V]) {
	e, s := m.set.FirstRest()
	m = Map[K, V]{s}

	if e == nil {
		return nil, nil, m
	}

	return &e.key, &e.value, m
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
