package hamt

// Dictionary represents a dictionary (associative array).
type Dictionary struct {
	Set
}

// NewDictionary creates a new dictionary.
func NewDictionary() Dictionary {
	return Dictionary{NewSet()}
}

// Insert inserts a key-value pair into a dictionary.
func (d Dictionary) Insert(k Entry, v interface{}) Dictionary {
	return Dictionary{d.Set.Insert(newKeyValue(k, v))}
}

// Delete deletes a pair of a key and a value from a dictionary.
func (d Dictionary) Delete(k Entry) Dictionary {
	return Dictionary{d.Set.Delete(k)}
}

// Find finds a value corresponding to a given key from a dictionary.
// It returns nil if no value is found.
func (d Dictionary) Find(k Entry) interface{} {
	e := d.Set.find(k)

	if e == nil {
		return nil
	}

	return e.(keyValue).value
}

// FirstRest returns a key-value pair in a dictionary and a rest of the dictionary.
// This method is useful for iteration.
// The key and value would be nil if the dictionary is empty.
func (d Dictionary) FirstRest() (Entry, interface{}, Dictionary) {
	e, s := d.Set.FirstRest()
	d = Dictionary{s}

	if e == nil {
		return nil, nil, d
	}

	kv := e.(keyValue)
	return kv.key, kv.value, d
}

// Size returns a size of a dictionary.
func (d Dictionary) Size() int {
	return d.Set.Size()
}
