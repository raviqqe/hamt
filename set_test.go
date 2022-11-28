package hamt

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	NewSet[entryInt]()
}

func TestSetInsert(t *testing.T) {
	s := NewSet[entryInt]()

	for i := 0; i < iterations; i++ {
		e := entryInt(rand.Int31())
		s = s.Insert(e)
		assert.True(t, s.Include(e))
	}
}

func TestSetOperations(t *testing.T) {
	s := NewSet[entryInt]()

	for i := 0; i < iterations; i++ {
		assert.Equal(t, s.hamt.Size(), s.Size())

		e := entryInt(rand.Int31() % 256)
		var ss Set[entryInt]

		if rand.Int()%2 == 0 {
			ss = s.Insert(e)

			assert.True(t, ss.Include(e))

			if s.Include(e) {
				assert.Equal(t, s.Size(), ss.Size())
			} else {
				assert.Equal(t, s.Size()+1, ss.Size())
			}
		} else {
			ss = s.Delete(e)

			assert.False(t, ss.Include(e))

			if s.Include(e) {
				assert.Equal(t, s.Size()-1, ss.Size())
			} else {
				assert.Equal(t, s.Size(), ss.Size())
			}
		}

		s = ss
	}
}

func TestSetFirstRest(t *testing.T) {
	s := NewSet[entryInt]()
	_, ss, ok := s.FirstRest()

	assert.False(t, ok)
	assert.Equal(t, 0, ss.Size())

	s = s.Insert(entryInt(42))
	e, ss, ok := s.FirstRest()

	assert.True(t, ok)
	assert.Equal(t, entryInt(42), e)
	assert.Equal(t, 0, ss.Size())

	s = s.Insert(entryInt(2049))
	size := s.Size()

	for i := 0; i < size; i++ {
		_, s, ok = s.FirstRest()

		assert.True(t, ok)
		assert.Equal(t, 1-i, s.Size())
	}
}

func TestSetForEach(t *testing.T) {
	s := NewSet[entryInt]()
	err := s.ForEach(func(entry entryInt) error {
		assert.Fail(t, "for-each callback called on empty set")
		return nil
	})
	assert.NoError(t, err)

	s = s.Insert(entryInt(42))
	entries := make([]entryInt, 0)
	err = s.ForEach(func(entry entryInt) error {
		entries = append(entries, entry)
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, []entryInt{42}, entries)

	s = s.Insert(entryInt(2049))
	entries = make([]entryInt, 0)
	err = s.ForEach(func(entry entryInt) error {
		entries = append(entries, entry)
		return nil
	})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []entryInt{42, 2049}, entries)
}

func TestSetMerge(t *testing.T) {
	for _, ss := range [][3]Set[entryInt]{
		{
			NewSet[entryInt](),
			NewSet[entryInt](),
			NewSet[entryInt](),
		},
		{
			NewSet[entryInt]().Insert(entryInt(1)),
			NewSet[entryInt](),
			NewSet[entryInt]().Insert(entryInt(1)),
		},
		{
			NewSet[entryInt](),
			NewSet[entryInt]().Insert(entryInt(1)),
			NewSet[entryInt]().Insert(entryInt(1)),
		},
		{
			NewSet[entryInt]().Insert(entryInt(2)),
			NewSet[entryInt]().Insert(entryInt(1)),
			NewSet[entryInt]().Insert(entryInt(1)).Insert(entryInt(2))},
	} {
		assert.Equal(t, ss[2], ss[0].Merge(ss[1]))
	}
}

func TestSetSize(t *testing.T) {
	assert.Equal(t, 0, NewSet[entryInt]().Size())
}

func BenchmarkSetInsert(b *testing.B) {
	s := NewSet[entryInt]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			s = s.Insert(entryInt(i))
		}
	}
}

func BenchmarkSetSize(b *testing.B) {
	s := NewSet[entryInt]()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < iterations; i++ {
			s = s.Insert(entryInt(i))
			b.Log(s.Size())
		}
	}
}
