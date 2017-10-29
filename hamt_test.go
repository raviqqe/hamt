package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHamt(t *testing.T) {
	NewHamt(0)
}

func TestHamtInsert(t *testing.T) {
	h := NewHamt(0).Insert(EntryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, EntryInt(42), h.Find(EntryInt(42)).(EntryInt))

	h = h.Insert(EntryInt(2049))

	assert.Equal(t, 2, h.Size())
	assert.Equal(t, EntryInt(42), h.Find(EntryInt(42)).(EntryInt))
	assert.Equal(t, EntryInt(2049), h.Find(EntryInt(2049)).(EntryInt))
}

func TestHamtDelete(t *testing.T) {
	h := NewHamt(0).Insert(EntryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, EntryInt(42), h.Find(EntryInt(42)).(EntryInt))

	h = h.Delete(EntryInt(42))

	assert.Equal(t, 0, h.Size())
	assert.Equal(t, nil, h.Find(EntryInt(42)))
}

func TestHamtFind(t *testing.T) {
	h := NewHamt(0)
	h.Find(EntryInt(42))
}

func TestHamtFirstRest(t *testing.T) {
	var n Node = NewHamt(0)
	e, m := n.FirstRest()

	assert.Equal(t, nil, e)
	assert.Equal(t, 0, m.Size())

	n = n.Insert(EntryInt(42))
	e, m = n.FirstRest()

	assert.Equal(t, EntryInt(42), e)
	assert.Equal(t, 0, m.Size())

	n = n.Insert(EntryInt(2049))
	s := n.Size()

	for i := 0; i < s; i++ {
		e, n = n.FirstRest()

		assert.NotEqual(t, nil, e)
		assert.Equal(t, 1-i, n.Size())
	}
}

func TestHamtSize(t *testing.T) {
	assert.Equal(t, 0, NewHamt(0).Size())
}

func TestArity(t *testing.T) {
	assert.Equal(t, arity, int(1<<arityBits))
}
