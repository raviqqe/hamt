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
	v, ok := b.Find(entryInt(42))
	assert.True(t, ok)
	assert.Equal(t, entryInt(42), v)

	b = b.Insert(entryInt(42))

	assert.Equal(t, 1, b.Size())
	v, ok = b.Find(entryInt(42))
	assert.True(t, ok)
	assert.Equal(t, entryInt(42), v)

	b = b.Insert(entryInt(2049))

	assert.Equal(t, 2, b.Size())
	v, ok = b.Find(entryInt(42))
	assert.True(t, ok)
	assert.Equal(t, entryInt(42), v)
	v, ok = b.Find(entryInt(2049))
	assert.True(t, ok)
	assert.Equal(t, entryInt(2049), v)

	b = b.Insert(entryInt(2049))

	assert.Equal(t, 2, b.Size())
	v, ok = b.Find(entryInt(42))
	assert.True(t, ok)
	assert.Equal(t, entryInt(42), v)
	v, ok = b.Find(entryInt(2049))
	assert.True(t, ok)
	assert.Equal(t, entryInt(2049), v)
}

func TestBucketInsertAsMap(t *testing.T) {
	kv := newTestKeyValue(0, "foo")
	b := newBucket[keyValue[entryInt, string]]().Insert(kv)

	assert.Equal(t, 1, b.Size())
	v, ok := b.Find(kv)
	assert.True(t, ok)
	assert.EqualValues(t, kv, v)

	new := newTestKeyValue(0, "bar")
	b = b.Insert(new)

	assert.Equal(t, 1, b.Size())
	v, ok = b.Find(kv)
	assert.True(t, ok)
	assert.EqualValues(t, new, v)
}

func TestBucketDelete(t *testing.T) {
	b, changed := newBucket[entryInt]().Insert(entryInt(42)).Delete(entryInt(42))

	assert.True(t, changed)
	assert.Equal(t, 0, b.Size())
	_, ok := b.Find(entryInt(42))
	assert.False(t, ok)
}

func TestBucketDeleteNonExistentEntries(t *testing.T) {
	b, changed := newBucket[entryInt]().Delete(entryInt(42))

	assert.False(t, changed)
	assert.Equal(t, 0, b.Size())

	b, changed = newBucket[entryInt]().Insert(entryInt(42)).Delete(entryInt(2049))

	assert.False(t, changed)
	assert.Equal(t, 1, b.Size())
	v, ok := b.Find(entryInt(42))
	assert.True(t, ok)
	assert.Equal(t, entryInt(42), v)
}

func TestBucketFind(t *testing.T) {
	_, ok := newBucket[entryInt]().Find(entryInt(42))
	assert.False(t, ok)
}

func TestBucketFirstRest(t *testing.T) {
	_, b, ok := newBucket[entryInt]().FirstRest()

	assert.False(t, ok)
	assert.Equal(t, 0, b.Size())

	b = b.Insert(entryInt(42))
	e, r, ok := b.FirstRest()

	assert.True(t, ok)
	assert.Equal(t, entryInt(42), e)
	assert.Equal(t, 0, r.Size())

	b = b.Insert(entryInt(2049))
	s := b.Size()

	for i := 0; i < s; i++ {
		_, b, ok = b.FirstRest()

		assert.True(t, ok)
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
