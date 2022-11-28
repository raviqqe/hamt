package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBucket(t *testing.T) {
	b := newBucket[entryInt]()

	assert.Equal(t, 0, b.Size())
}

func TestBucketInsert(t *testing.T) {
	b := newBucket[entryInt]().Insert(entryInt(42))

	assert.Equal(t, 1, b.Size())
	assert.Equal(t, entryInt(42), *b.Find(entryInt(42)))

	b = b.Insert(entryInt(42))

	assert.Equal(t, 1, b.Size())
	assert.Equal(t, entryInt(42), *b.Find(entryInt(42)))

	b = b.Insert(entryInt(2049))

	assert.Equal(t, 2, b.Size())
	assert.Equal(t, entryInt(42), *b.Find(entryInt(42)))
	assert.Equal(t, entryInt(2049), *b.Find(entryInt(2049)))

	b = b.Insert(entryInt(2049))

	assert.Equal(t, 2, b.Size())
	assert.Equal(t, entryInt(42), *b.Find(entryInt(42)))
	assert.Equal(t, entryInt(2049), *b.Find(entryInt(2049)))
}

func TestBucketInsertAsMap(t *testing.T) {
	kv := newTestKeyValue(0, "foo")
	b := newBucket[keyValue[entryInt, string]]().Insert(kv)

	assert.Equal(t, 1, b.Size())
	assert.EqualValues(t, kv, *b.Find(kv))

	new := newTestKeyValue(0, "bar")
	b = b.Insert(new)

	assert.Equal(t, 1, b.Size())
	assert.EqualValues(t, new, *b.Find(kv))
}

func TestBucketDelete(t *testing.T) {
	b, changed := newBucket[entryInt]().Insert(entryInt(42)).Delete(entryInt(42))

	assert.True(t, changed)
	assert.Equal(t, 0, b.Size())
	assert.Nil(t, b.Find(entryInt(42)))
}

func TestBucketDeleteNonExistentEntries(t *testing.T) {
	b, changed := newBucket[entryInt]().Delete(entryInt(42))

	assert.False(t, changed)
	assert.Equal(t, 0, b.Size())

	b, changed = newBucket[entryInt]().Insert(entryInt(42)).Delete(entryInt(2049))

	assert.False(t, changed)
	assert.Equal(t, 1, b.Size())
	assert.Equal(t, entryInt(42), *b.Find(entryInt(42)))
}

func TestBucketFind(t *testing.T) {
	assert.Nil(t, newBucket[entryInt]().Find(entryInt(42)))
}

func TestBucketFirstRest(t *testing.T) {
	e, b := newBucket[entryInt]().FirstRest()

	assert.Nil(t, e)
	assert.Equal(t, 0, b.Size())

	b = b.Insert(entryInt(42))
	e, r := b.FirstRest()

	assert.Equal(t, entryInt(42), *e)
	assert.Equal(t, 0, r.Size())

	b = b.Insert(entryInt(2049))
	s := b.Size()

	for i := 0; i < s; i++ {
		e, b = b.FirstRest()

		assert.NotEqual(t, nil, e)
		assert.Equal(t, 1-i, b.Size())
	}
}

func TestBucketState(t *testing.T) {
	var b node[entryInt] = newBucket[entryInt]()

	assert.Equal(t, empty, b.State())

	b = b.Insert(entryInt(42))

	assert.Equal(t, singleton, b.State())

	b = b.Insert(entryInt(2049))

	assert.Equal(t, more, b.State())

	b = b.Insert(entryInt(0))

	assert.Equal(t, more, b.State())
}
