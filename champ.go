package champ

// Champ represents a CHAMP data structure.
type Champ struct{}

// NewChamp creates a new CHAMP.
func NewChamp() Champ {
	return Champ{}
}

// Insert inserts a key into a CHAMP.
func (champ Champ) Insert(x interface{}) Champ {
	return champ
}

// Delete deletes a key from a CHAMP.
func (champ Champ) Delete(x interface{}) Champ {
	return champ
}

// Find finds a key in a CHAMP.
func (champ Champ) Find(x interface{}) interface{} {
	return nil
}

// FirstRest returns a first key and a CHAMP without it.
func (champ Champ) FirstRest() (interface{}, Champ) {
	return nil, champ
}
