package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestKeyValue(k uint32, v string) Entry {
	return newKeyValue(EntryInt(k), v)
}

func TestNewKeyValue(t *testing.T) {
	newKeyValue(EntryInt(42), "value")
}

func TestKeyValueAsEntry(t *testing.T) {
	t.Log(Entry(newKeyValue(EntryInt(42), "value")))
}

func TestKeyValueHash(t *testing.T) {
	newKeyValue(EntryInt(42), "value").Hash()
}

func TestKeyValueEqual(t *testing.T) {
	k := EntryInt(42)
	kv := newKeyValue(k, "value")

	assert.True(t, kv.Equal(kv))
	assert.True(t, kv.Equal(k))
	assert.False(t, kv.Equal(newKeyValue(EntryInt(2049), "value")))
	assert.False(t, kv.Equal(EntryInt(2049)))
}
