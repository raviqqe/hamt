package hamt

// Map represents a map (associative array).
type Map struct {
	set Set
}

// NewMap creates a new map.
func NewMap() Map {
	return Map{NewSet()}
}

// Insert inserts a key-value pair into a map.
func (m Map) Insert(k Entry, v interface{}) Map {
	return Map{m.set.Insert(newKeyValue(k, v))}
}

// Delete deletes a pair of a key and a value from a map.
func (m Map) Delete(k Entry) Map {
	return Map{m.set.Delete(k)}
}

// Find finds a value corresponding to a given key from a map.
// It returns nil if no value is found.
func (m Map) Find(k Entry) interface{} {
	e := m.set.find(k)

	if e == nil {
		return nil
	}

	return e.(keyValue).value
}

// Include returns true if a key-value pair corresponding with a given key is
// included in a map, or false otherwise.
func (m Map) Include(k Entry) bool {
	return m.Find(k) != nil
}

// FirstRest returns a key-value pair in a map and a rest of the map.
// This method is useful for iteration.
// The key and value would be nil if the map is empty.
func (m Map) FirstRest() (Entry, interface{}, Map) {
	e, s := m.set.FirstRest()
	m = Map{s}

	if e == nil {
		return nil, nil, m
	}

	kv := e.(keyValue)
	return kv.key, kv.value, m
}

// Size returns a size of a map.
func (m Map) Size() int {
	return m.set.Size()
}
