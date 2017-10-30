package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestKeyValue(k uint32, v string) Entry {
	return newKeyValue(entryInt(k), v)
}

func TestNewKeyValue(t *testing.T) {
	newKeyValue(entryInt(42), "value")
}

func TestKeyValueAsEntry(t *testing.T) {
	t.Log(Entry(newKeyValue(entryInt(42), "value")))
}

func TestKeyValueHash(t *testing.T) {
	newKeyValue(entryInt(42), "value").Hash()
}

func TestKeyValueEqual(t *testing.T) {
	k := entryInt(42)
	kv := newKeyValue(k, "value")

	assert.True(t, kv.Equal(kv))
	assert.True(t, kv.Equal(k))
	assert.False(t, kv.Equal(newKeyValue(entryInt(2049), "value")))
	assert.False(t, kv.Equal(entryInt(2049)))
}
