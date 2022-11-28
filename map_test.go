package hamt

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	NewMap[entryInt, string]()
}

func TestMapInsert(t *testing.T) {
	m := NewMap[entryInt, string]()

	for i := 0; i < iterations; i++ {
		e := entryInt(rand.Int31())
		m = m.Insert(e, "value")
		v, ok := m.Find(e)
		assert.True(t, ok)
		assert.Equal(t, "value", v)
	}
}

func TestMapOperations(t *testing.T) {
	m := NewMap[entryInt, string]()

	for i := 0; i < iterations; i++ {
		k := entryInt(rand.Int31() % 256)
		var mm Map[entryInt, string]

		if rand.Int()%2 == 0 {
			mm = m.Insert(k, "value")

			v, ok := mm.Find(k)
			assert.True(t, ok)
			assert.Equal(t, "value", v)

			if m.Include(k) {
				assert.Equal(t, m.Size(), mm.Size())
			} else {
				assert.Equal(t, m.Size()+1, mm.Size())
			}
		} else {
			mm = m.Delete(k)

			_, ok := mm.Find(k)
			assert.False(t, ok)

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
	m := NewMap[entryInt, string]()
	_, _, mm, ok := m.FirstRest()

	assert.False(t, ok)
	assert.Equal(t, 0, mm.Size())

	m = m.Insert(entryInt(42), "value")
	k, v, mm, ok := m.FirstRest()

	assert.True(t, ok)
	assert.Equal(t, entryInt(42), k)
	assert.Equal(t, "value", v)
	assert.Equal(t, 0, mm.Size())

	m = m.Insert(entryInt(2049), "value")
	s := m.Size()

	for i := 0; i < s; i++ {
		_, v, m, ok = m.FirstRest()

		assert.True(t, ok)
		assert.Equal(t, "value", v)
		assert.Equal(t, 1-i, m.Size())
	}
}

func TestMapForEach(t *testing.T) {
	m := NewMap[entryInt, string]()
	err := m.ForEach(func(key entryInt, val string) error {
		assert.Fail(t, "for-each callback called on empty set")
		return nil
	})
	assert.NoError(t, err)

	m = m.Insert(entryInt(42), "value")
	kvs := make([]keyValue[entryInt, string], 0)
	want := []keyValue[entryInt, string]{
		{
			key:   entryInt(42),
			value: "value",
		},
	}
	err = m.ForEach(func(key entryInt, val string) error {
		kvs = append(kvs, keyValue[entryInt, string]{
			key:   key,
			value: val,
		})
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, want, kvs)

	m = m.Insert(entryInt(2049), "value2")
	kvs = make([]keyValue[entryInt, string], 0)
	want = []keyValue[entryInt, string]{
		{
			key:   entryInt(42),
			value: "value",
		},
		{
			key:   entryInt(2049),
			value: "value2",
		},
	}
	err = m.ForEach(func(key entryInt, val string) error {
		kvs = append(kvs, keyValue[entryInt, string]{
			key:   key,
			value: val,
		})
		return nil
	})
	assert.NoError(t, err)
	assert.ElementsMatch(t, want, kvs)
}

func TestMapMerge(t *testing.T) {
	for _, ms := range [][3]Map[entryInt, string]{
		{
			NewMap[entryInt, string](),
			NewMap[entryInt, string](),
			NewMap[entryInt, string](),
		},
		{
			NewMap[entryInt, string]().Insert(entryInt(1), "foo"),
			NewMap[entryInt, string](),
			NewMap[entryInt, string]().Insert(entryInt(1), "foo"),
		},
		{
			NewMap[entryInt, string](),
			NewMap[entryInt, string]().Insert(entryInt(1), "foo"),
			NewMap[entryInt, string]().Insert(entryInt(1), "foo"),
		},
		{
			NewMap[entryInt, string]().Insert(entryInt(2), "foo"),
			NewMap[entryInt, string]().Insert(entryInt(1), "foo"),
			NewMap[entryInt, string]().Insert(entryInt(1), "foo").Insert(entryInt(2), "foo")},
	} {
		assert.Equal(t, ms[2], ms[0].Merge(ms[1]))
	}
}

func TestMapSize(t *testing.T) {
	assert.Equal(t, 0, NewMap[entryInt, string]().Size())
}
