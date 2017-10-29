package champ

// Champ represents a CHAMP data structure.
type Champ struct{}

// NewChamp creates a new CHAMP.
func NewChamp() Champ {
	return Champ{}
}

// Insert inserts a value into a CHAMP.
func (champ Champ) Insert(e Entry) Champ {
	return champ
}

// Delete deletes a value from a CHAMP.
func (champ Champ) Delete(e Entry) Champ {
	return champ
}

// Find finds a value in a CHAMP.
func (champ Champ) Find(e Entry) Entry {
	return nil
}

// FirstRest returns a first value and a CHAMP without it.
func (champ Champ) FirstRest() (Entry, Champ) {
	return nil, champ
}

// Size returns a size of a CHAMP.
func (champ Champ) Size() int {
	return 0
}
