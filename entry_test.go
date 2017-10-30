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

type EntryKeyValue struct {
	key   uint32
	value string
}

func NewEntryKeyValue(k uint32, v string) EntryKeyValue {
	return EntryKeyValue{k, v}
}

func (kv EntryKeyValue) Hash() uint32 {
	return kv.key
}

func (kv EntryKeyValue) Equal(e Entry) bool {
	x, ok := e.(EntryKeyValue)

	if !ok {
		return false
	}

	return kv.key == x.key
}

func TestEntry(t *testing.T) {
	t.Log(Entry(EntryInt(42)))
}

func TestEntryKey(t *testing.T) {
	assert.Equal(t, uint32(42), Entry(EntryInt(42)).Hash())
}
