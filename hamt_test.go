package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHamt(t *testing.T) {
	newHamt(0)
}

func TestHamtInsert(t *testing.T) {
	h := newHamt(0).Insert(EntryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, EntryInt(42), h.Find(EntryInt(42)).(EntryInt))

	h = h.Insert(EntryInt(2049))

	assert.Equal(t, 2, h.Size())
	assert.Equal(t, EntryInt(42), h.Find(EntryInt(42)).(EntryInt))
	assert.Equal(t, EntryInt(2049), h.Find(EntryInt(2049)).(EntryInt))
}

func TestHamtDelete(t *testing.T) {
	h := newHamt(0).Insert(EntryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, EntryInt(42), h.Find(EntryInt(42)).(EntryInt))

	h, changed := h.Delete(EntryInt(42))

	assert.True(t, changed)
	assert.Equal(t, 0, h.Size())
	assert.Equal(t, nil, h.Find(EntryInt(42)))
}

func TestHamtDeleteWithManyEntries(t *testing.T) {
	var h node = newHamt(0)

	for i := 0; i < 10000; i++ {
		h = h.Insert(EntryInt(int32(i)))
	}

	h, _ = h.Delete(EntryInt(42))

	assert.Equal(t, nil, h.Find(EntryInt(42)))
}

func TestHamtFind(t *testing.T) {
	assert.Equal(t, nil, newHamt(0).Find(EntryInt(42)))
}

func TestHamtFirstRest(t *testing.T) {
	var n node = newHamt(0)
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
	assert.Equal(t, 0, newHamt(0).Size())
}

func TestArity(t *testing.T) {
	assert.Equal(t, arity, int(1<<arityBits))
}
