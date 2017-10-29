package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHamt(t *testing.T) {
	NewHamt(0)
}

func TestHamtInsert(t *testing.T) {
	hamt := NewHamt(0)
	hamt.Insert(EntryInt(42))
}

func TestHamtDelete(t *testing.T) {
	hamt := NewHamt(0)
	hamt.Delete(EntryInt(42))
}

func TestHamtFind(t *testing.T) {
	hamt := NewHamt(0)
	hamt.Find(EntryInt(42))
}

func TestHamtFirstRest(t *testing.T) {
	hamt := NewHamt(0)
	hamt.FirstRest()
}

func TestHamtSize(t *testing.T) {
	hamt := NewHamt(0)
	assert.Equal(t, 0, hamt.Size())
}

func TestArity(t *testing.T) {
	assert.Equal(t, arity, int(1<<arityBits))
}
