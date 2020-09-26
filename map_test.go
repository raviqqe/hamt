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
		e := entryInt(rand.Int31())
		m = m.Insert(e, "value")
		assert.Equal(t, "value", m.Find(e))
	}
}

func TestMapOperations(t *testing.T) {
	m := NewMap()

	for i := 0; i < iterations; i++ {
		k := entryInt(rand.Int31() % 256)
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

	m = m.Insert(entryInt(42), "value")
	k, v, mm = m.FirstRest()

	assert.Equal(t, entryInt(42), k)
	assert.Equal(t, "value", v)
	assert.Equal(t, 0, mm.Size())

	m = m.Insert(entryInt(2049), "value")
	s := m.Size()

	for i := 0; i < s; i++ {
		k, v, m = m.FirstRest()

		assert.NotEqual(t, nil, k)
		assert.Equal(t, "value", v)
		assert.Equal(t, 1-i, m.Size())
	}
}

func TestMapForEach(t *testing.T) {
	m := NewMap()
	err := m.ForEach(func(key Entry, val interface{}) error {
		assert.Fail(t, "for-each callback called on empty set")
		return nil
	})
	assert.NoError(t, err)

	m = m.Insert(entryInt(42), "value")
	kvs := make([]keyValue, 0)
	want := []keyValue{
		{
			key:   entryInt(42),
			value: "value",
		},
	}
	err = m.ForEach(func(key Entry, val interface{}) error {
		kvs = append(kvs, keyValue{
			key:   key,
			value: val,
		})
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, want, kvs)

	m = m.Insert(entryInt(2049), "value2")
	kvs = make([]keyValue, 0)
	want = []keyValue{
		{
			key:   entryInt(42),
			value: "value",
		},
		{
			key:   entryInt(2049),
			value: "value2",
		},
	}
	err = m.ForEach(func(key Entry, val interface{}) error {
		kvs = append(kvs, keyValue{
			key:   key,
			value: val,
		})
		return nil
	})
	assert.NoError(t, err)
	assert.ElementsMatch(t, want, kvs)
}

func TestMapMerge(t *testing.T) {
	for _, ms := range [][3]Map{
		{
			NewMap(),
			NewMap(),
			NewMap(),
		},
		{
			NewMap().Insert(entryInt(1), "foo"),
			NewMap(),
			NewMap().Insert(entryInt(1), "foo"),
		},
		{
			NewMap(),
			NewMap().Insert(entryInt(1), "foo"),
			NewMap().Insert(entryInt(1), "foo"),
		},
		{
			NewMap().Insert(entryInt(2), "foo"),
			NewMap().Insert(entryInt(1), "foo"),
			NewMap().Insert(entryInt(1), "foo").Insert(entryInt(2), "foo")},
	} {
		assert.Equal(t, ms[2], ms[0].Merge(ms[1]))
	}
}

func TestMapSize(t *testing.T) {
	assert.Equal(t, 0, NewMap().Size())
}
