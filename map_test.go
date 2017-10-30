package hamt

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	NewMap()
}

func TestMapInsert(t *testing.T) {
	m := NewMap()

	for i := 0; i < iterations; i++ {
		e := EntryInt(rand.Int31())
		m = m.Insert(e, "value")
		assert.Equal(t, "value", m.Find(e))
	}
}

func TestMapOperations(t *testing.T) {
	m := NewMap()

	for i := 0; i < iterations; i++ {
		k := EntryInt(rand.Int31() % 256)
		var mm Map

		if rand.Int()%2 == 0 {
			mm = m.Insert(k, "value")

			assert.Equal(t, "value", mm.Find(k))

			if m.Include(k) {
				assert.Equal(t, m.Size(), mm.Size())
			} else {
				assert.Equal(t, m.Size()+1, mm.Size())
			}
		} else {
			mm = m.Delete(k)

			assert.Equal(t, nil, mm.Find(k))

			if m.Include(k) {
				assert.Equal(t, m.Size()-1, mm.Size())
			} else {
				assert.Equal(t, m.Size(), mm.Size())
			}
		}

		m = mm
	}
}

func TestMapFirstRest(t *testing.T) {
	m := NewMap()
	k, v, mm := m.FirstRest()

	assert.Equal(t, nil, k)
	assert.Equal(t, nil, v)
	assert.Equal(t, 0, mm.Size())

	m = m.Insert(EntryInt(42), "value")
	k, v, mm = m.FirstRest()

	assert.Equal(t, EntryInt(42), k)
	assert.Equal(t, "value", v)
	assert.Equal(t, 0, mm.Size())

	m = m.Insert(EntryInt(2049), "value")
	s := m.Size()

	for i := 0; i < s; i++ {
		k, v, m = m.FirstRest()

		assert.NotEqual(t, nil, k)
		assert.Equal(t, "value", v)
		assert.Equal(t, 1-i, m.Size())
	}
}

func TestMapSize(t *testing.T) {
	assert.Equal(t, 0, NewMap().Size())
}
