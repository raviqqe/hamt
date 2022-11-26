package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entryInt uint32

func (i entryInt) Hash() uint32 {
	return uint32(i)
}

func (i entryInt) Equal(j entryInt) bool {
	return i == j
}

func TestEntry(t *testing.T) {
	t.Log(Entry[entryInt](entryInt(42)))
}

func TestEntryKey(t *testing.T) {
	assert.Equal(t, uint32(42), entryInt(42).Hash())
}
