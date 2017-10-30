package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHamt(t *testing.T) {
	newHamt(0)
}

func TestHamtInsert(t *testing.T) {
	h := newHamt(0).Insert(entryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, entryInt(42), h.Find(entryInt(42)).(entryInt))

	h = h.Insert(entryInt(2049))

	assert.Equal(t, 2, h.Size())
	assert.Equal(t, entryInt(42), h.Find(entryInt(42)).(entryInt))
	assert.Equal(t, entryInt(2049), h.Find(entryInt(2049)).(entryInt))
}

func TestHamtInsertAsMap(t *testing.T) {
	kv := newTestKeyValue(0, "foo")
	h := newHamt(0).Insert(kv)

	assert.Equal(t, 1, h.Size())
	assert.EqualValues(t, kv, h.Find(kv))

	new := newTestKeyValue(0, "bar")
	h = h.Insert(new)

	assert.Equal(t, 1, h.Size())
	assert.EqualValues(t, new, h.Find(kv))
}

func TestHamtInsertWithBucketCreation(t *testing.T) {
	h := newHamt(7).Insert(entryInt(0)).Insert(entryInt(0x80000000))

	b, ok := h.(hamt).children[0].(bucket)

	assert.True(t, ok)
	assert.Equal(t, 2, b.Size())
}

func TestHamtDelete(t *testing.T) {
	h := newHamt(0).Insert(entryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, entryInt(42), h.Find(entryInt(42)).(entryInt))

	h, changed := h.Delete(entryInt(42))

	assert.True(t, changed)
	assert.Equal(t, 0, h.Size())
	assert.Equal(t, nil, h.Find(entryInt(42)))
}

func TestHamtDeleteWithManyEntries(t *testing.T) {
	var h node = newHamt(0)

	for i := 0; i < iterations; i++ {
		h = h.Insert(entryInt(uint32(i)))
	}

	assert.Equal(t, iterations, h.Size())

	for i := 0; i < iterations; i++ {
		e := entryInt(uint32(i))
		g, ok := h.Delete(e)

		assert.True(t, ok)
		assert.Nil(t, g.Find(e))
		assert.Equal(t, h.Size()-1, g.Size())

		h = g
	}

	assert.Equal(t, 0, h.Size())
}

func TestHamtFind(t *testing.T) {
	assert.Equal(t, nil, newHamt(0).Find(entryInt(42)))
}

func TestHamtFirstRest(t *testing.T) {
	var n node = newHamt(0)
	e, m := n.FirstRest()

	assert.Equal(t, nil, e)
	assert.Equal(t, 0, m.Size())

	n = n.Insert(entryInt(42))
	e, m = n.FirstRest()

	assert.Equal(t, entryInt(42), e)
	assert.Equal(t, 0, m.Size())

	n = n.Insert(entryInt(2049))
	s := n.Size()

	for i := 0; i < s; i++ {
		e, n = n.FirstRest()

		assert.NotEqual(t, nil, e)
		assert.Equal(t, 1-i, n.Size())
	}
}

func TestHamtFirstRestWithManyEntries(t *testing.T) {
	var h node = newHamt(0)

	for i := 0; i < iterations; i++ {
		h = h.Insert(entryInt(uint32(i)))
	}

	assert.Equal(t, iterations, h.Size())

	for i := 0; i < iterations; i++ {
		e, g := h.FirstRest()

		assert.NotNil(t, e)
		assert.Equal(t, h.Size()-1, g.Size())

		h = g
	}
}

func TestHamtSize(t *testing.T) {
	assert.Equal(t, 0, newHamt(0).Size())
}

func TestHamtCalculateIndex(t *testing.T) {
	e := entryInt(0xffffffff)

	for i := 0; i < 6; i++ {
		assert.Equal(t, 0x1f, newHamt(uint8(i)).calculateIndex(e))
	}

	assert.Equal(t, 3, newHamt(6).calculateIndex(e))
}

func TestArity(t *testing.T) {
	assert.Equal(t, arity, int(1<<arityBits))
}

func BenchmarkHamtInsert(b *testing.B) {
	var h node = newHamt(0)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			h = h.Insert(entryInt(i % (iterations / 3)))
		}
	}
}

func BenchmarkHamtDelete(b *testing.B) {
	var h node = newHamt(0)

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			h = h.Insert(entryInt(i))
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			h, _ = h.Delete(entryInt(i))
		}
	}
}
