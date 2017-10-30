package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EntryInt uint32

func (i EntryInt) Hash() uint32 {
	return uint32(i)
}

func (i EntryInt) Equal(e Entry) bool {
	j, ok := e.(EntryInt)

	if !ok {
		return false
	}

	return i == j
}

func TestEntry(t *testing.T) {
	t.Log(Entry(EntryInt(42)))
}

func TestEntryKey(t *testing.T) {
	assert.Equal(t, uint32(42), Entry(EntryInt(42)).Hash())
}
