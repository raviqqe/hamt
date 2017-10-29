package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBucket(t *testing.T) {
	b := newBucket(nil)

	assert.Equal(t, 0, b.Size())
}

func TestBucketInsert(t *testing.T) {
	b := newBucket(nil).Insert(EntryInt(42))

	assert.Equal(t, 1, b.Size())
	assert.Equal(t, EntryInt(42), b.Find(EntryInt(42)))
}

func TestBucketDelete(t *testing.T) {
	b := newBucket(nil).Insert(EntryInt(42)).Delete(EntryInt(42))

	assert.Equal(t, 0, b.Size())
	assert.Equal(t, nil, b.Find(EntryInt(42)))
}

func TestBucketFind(t *testing.T) {
	assert.Equal(t, nil, newBucket(nil).Find(EntryInt(42)))
}

func TestBucketFirstRest(t *testing.T) {
	e, b := newBucket(nil).FirstRest()

	assert.Equal(t, nil, e)
	assert.Equal(t, 0, b.Size())

	b = b.Insert(EntryInt(42))
	e, r := b.FirstRest()

	assert.Equal(t, EntryInt(42), e)
	assert.Equal(t, 0, r.Size())

	b = b.Insert(EntryInt(2049))
	s := b.Size()

	for i := 0; i < s; i++ {
		e, b = b.FirstRest()

		assert.NotEqual(t, nil, e)
		assert.Equal(t, 1-i, b.Size())
	}
}
