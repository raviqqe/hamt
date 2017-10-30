package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entryInt uint32

func (i entryInt) Hash() uint32 {
	return uint32(i)
}

func (i entryInt) Equal(e Entry) bool {
	j, ok := e.(entryInt)

	if !ok {
		return false
	}

	return i == j
}

func TestEntry(t *testing.T) {
	t.Log(Entry(entryInt(42)))
}

func TestEntryKey(t *testing.T) {
	assert.Equal(t, uint32(42), Entry(entryInt(42)).Hash())
}
