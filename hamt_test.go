package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHamt(t *testing.T) {
	assert.Equal(t, newHamt[entryInt](0).Size(), 0)
}

func TestHamtInsert(t *testing.T) {
	h := newHamt[entryInt](0).Insert(entryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, entryInt(42), *h.Find(entryInt(42)))

	h = h.Insert(entryInt(2049))

	assert.Equal(t, 2, h.Size())
	assert.Equal(t, entryInt(42), *h.Find(entryInt(42)))
	assert.Equal(t, entryInt(2049), *h.Find(entryInt(2049)))
}

func TestHamtInsertAsMap(t *testing.T) {
	kv := newTestKeyValue(0, "foo")
	h := newHamt[keyValue[entryInt, string]](0).Insert(kv)

	assert.Equal(t, 1, h.Size())
	assert.EqualValues(t, kv, *h.Find(kv))

	new := newTestKeyValue(0, "bar")
	h = h.Insert(new)

	assert.Equal(t, 1, h.Size())
	assert.EqualValues(t, new, *h.Find(kv))
}

func TestHamtInsertWithBucketCreation(t *testing.T) {
	h := newHamt[entryInt](7).Insert(entryInt(0)).Insert(entryInt(0x80000000))

	b, ok := h.(hamt[entryInt]).children[0].(bucket[entryInt])

	assert.True(t, ok)
	assert.Equal(t, 2, b.Size())
}

func TestHamtDelete(t *testing.T) {
	h := newHamt[entryInt](0).Insert(entryInt(42))

	assert.Equal(t, 1, h.Size())
	assert.Equal(t, entryInt(42), *h.Find(entryInt(42)))

	h, changed := h.Delete(entryInt(42))

	assert.True(t, changed)
	assert.Equal(t, 0, h.Size())
	assert.Nil(t, h.Find(entryInt(42)))
}

func TestHamtDeleteWithManyEntries(t *testing.T) {
	var h node[entryInt] = newHamt[entryInt](0)

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

func TestHamtDeletePanicWithUnnormalizedTree(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()

	e := entryInt(42)
	h := newHamt[entryInt](0)

	for i := range h.children {
		g := hamt[entryInt]{1, [32]any{}}

		if i == h.calculateIndex(e) {
			g.children[g.calculateIndex(e)] = e
		}

		h.children[i] = g
	}

	h.Delete(e)
}

func TestHamtFind(t *testing.T) {
	assert.Nil(t, newHamt[entryInt](0).Find(entryInt(42)))
}

func TestHamtFirstRest(t *testing.T) {
	var n node[entryInt] = newHamt[entryInt](0)
	e, m := n.FirstRest()

	assert.Nil(t, e)
	assert.Equal(t, 0, m.Size())

	n = n.Insert(entryInt(42))
	e, m = n.FirstRest()

	assert.Equal(t, entryInt(42), *e)
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
	var h node[entryInt] = newHamt[entryInt](0)

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

	assert.Equal(t, 0, h.Size())
}

func TestHamtForEach(t *testing.T) {
	var n node[entryInt] = newHamt[entryInt](0)
	err := n.ForEach(func(entry entryInt) error {
		assert.Fail(t, "for-each callback called on empty hamt")
		return nil
	})
	assert.NoError(t, err)

	n = n.Insert(entryInt(42))
	entries := make([]entryInt, 0)
	err = n.ForEach(func(entry entryInt) error {
		entries = append(entries, entry)
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, []entryInt{42}, entries)

	n = n.Insert(entryInt(2049))
	entries = make([]entryInt, 0)
	err = n.ForEach(func(entry entryInt) error {
		entries = append(entries, entry)
		return nil
	})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []entryInt{42, 2049}, entries)
}

func TestHamtForEachWithManyEntries(t *testing.T) {
	var h node[entryInt] = newHamt[entryInt](0)

	want := make([]entryInt, 0)
	for i := 0; i < iterations; i++ {
		e := entryInt(uint32(i))
		h = h.Insert(e)
		want = append(want, e)
	}

	assert.Equal(t, iterations, h.Size())

	entries := make([]entryInt, 0)
	err := h.ForEach(func(entry entryInt) error {
		entries = append(entries, entry)
		return nil
	})
	assert.NoError(t, err)
	assert.Len(t, entries, iterations)
	assert.ElementsMatch(t, want, entries)
}

func TestHamtState(t *testing.T) {
	var h node[entryInt] = newHamt[entryInt](0)

	assert.Equal(t, empty, h.State())

	h = h.Insert(entryInt(42))

	assert.Equal(t, singleton, h.State())

	h = h.Insert(entryInt(2049))

	assert.Equal(t, more, h.State())

	h = h.Insert(entryInt(0))

	assert.Equal(t, more, h.State())
}

func TestHamtSize(t *testing.T) {
	assert.Equal(t, 0, newHamt[entryInt](0).Size())
}

func TestHamtCalculateIndex(t *testing.T) {
	e := entryInt(0xffffffff)

	for i := 0; i < 6; i++ {
		assert.Equal(t, 0x1f, newHamt[entryInt](uint8(i)).calculateIndex(e))
	}

	assert.Equal(t, 3, newHamt[entryInt](6).calculateIndex(e))
}

func TestArity(t *testing.T) {
	assert.Equal(t, arity, int(1<<arityBits))
}

func BenchmarkHamtInsert(b *testing.B) {
	var h node[entryInt] = newHamt[entryInt](0)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			h = h.Insert(entryInt(i % (iterations / 3)))
		}
	}
}

func BenchmarkHamtDelete(b *testing.B) {
	var h node[entryInt] = newHamt[entryInt](0)

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

func BenchmarkHamtFirstRestIteration(b *testing.B) {
	var h node[entryInt] = newHamt[entryInt](0)
	for i := 0; i < iterations; i++ {
		h = h.Insert(entryInt(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hh := h
		for hh.Size() > 0 {
			_, hh = hh.FirstRest()
		}
	}
}

func BenchmarkHamtForEachIteration(b *testing.B) {
	var h node[entryInt] = newHamt[entryInt](0)
	for i := 0; i < iterations; i++ {
		h = h.Insert(entryInt(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = h.ForEach(func(entry entryInt) error {
			return nil
		})
	}
}

func BenchmarkBuiltinMapForEach(b *testing.B) {
	m := make(map[entryInt]struct{})
	for i := 0; i < iterations; i++ {
		m[entryInt(i)] = struct{}{}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for e := range m {
			_ = e
		}
	}
}

func BenchmarkBuiltinSliceForEach(b *testing.B) {
	m := make([]entryInt, 0)
	for i := 0; i < iterations; i++ {
		m = append(m, entryInt(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, e := range m {
			_ = e
		}
	}
}
