package hamt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHamt(t *testing.T) {
	NewHamt()
}

func TestHamtInsert(t *testing.T) {
	hamt := NewHamt()
	hamt.Insert(EntryInt(42))
}

func TestHamtDelete(t *testing.T) {
	hamt := NewHamt()
	hamt.Delete(EntryInt(42))
}

func TestHamtFind(t *testing.T) {
	hamt := NewHamt()
	hamt.Find(EntryInt(42))
}

func TestHamtFirstRest(t *testing.T) {
	hamt := NewHamt()
	hamt.FirstRest()
}

func TestHamtSize(t *testing.T) {
	hamt := NewHamt()
	assert.Equal(t, 0, hamt.Size())
}
