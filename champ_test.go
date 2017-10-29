package champ

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChamp(t *testing.T) {
	NewChamp()
}

func TestChampInsert(t *testing.T) {
	champ := NewChamp()
	champ.Insert(EntryInt(42))
}

func TestChampDelete(t *testing.T) {
	champ := NewChamp()
	champ.Delete(EntryInt(42))
}

func TestChampFind(t *testing.T) {
	champ := NewChamp()
	champ.Find(EntryInt(42))
}

func TestChampFirstRest(t *testing.T) {
	champ := NewChamp()
	champ.FirstRest()
}

func TestChampSize(t *testing.T) {
	champ := NewChamp()
	assert.Equal(t, 0, champ.Size())
}
