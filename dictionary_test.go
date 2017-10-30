package hamt

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDictionary(t *testing.T) {
	NewDictionary()
}

func TestDictionaryInsert(t *testing.T) {
	s := NewDictionary()

	for i := 0; i < iterations; i++ {
		e := EntryInt(rand.Int31())
		s = s.Insert(e, "value")
		assert.Equal(t, "value", s.Find(e))
	}
}

func TestDictionaryOperations(t *testing.T) {
	d := NewDictionary()

	for i := 0; i < iterations; i++ {
		k := EntryInt(rand.Int31() % 256)
		var dd Dictionary

		if rand.Int()%2 == 0 {
			dd = d.Insert(k, "value")

			assert.Equal(t, "value", dd.Find(k))

			if d.Include(k) {
				assert.Equal(t, d.Size(), dd.Size())
			} else {
				assert.Equal(t, d.Size()+1, dd.Size())
			}
		} else {
			dd = d.Delete(k)

			assert.Equal(t, nil, dd.Find(k))

			if d.Include(k) {
				assert.Equal(t, d.Size()-1, dd.Size())
			} else {
				assert.Equal(t, d.Size(), dd.Size())
			}
		}

		d = dd
	}
}

func TestDictionaryFirstRest(t *testing.T) {
	d := NewDictionary()
	k, v, dd := d.FirstRest()

	assert.Equal(t, nil, k)
	assert.Equal(t, nil, v)
	assert.Equal(t, 0, dd.Size())

	d = d.Insert(EntryInt(42), "value")
	k, v, dd = d.FirstRest()

	assert.Equal(t, EntryInt(42), k)
	assert.Equal(t, "value", v)
	assert.Equal(t, 0, dd.Size())

	d = d.Insert(EntryInt(2049), "value")
	s := d.Size()

	for i := 0; i < s; i++ {
		k, v, d = d.FirstRest()

		assert.NotEqual(t, nil, k)
		assert.Equal(t, "value", v)
		assert.Equal(t, 1-i, d.Size())
	}
}

func TestDictionarySize(t *testing.T) {
	assert.Equal(t, 0, NewDictionary().Size())
}
