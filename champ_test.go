package champ

import "testing"

func TestNewChamp(t *testing.T) {
	NewChamp()
}

func TestChampInsert(t *testing.T) {
	champ := NewChamp()
	champ.Insert(123)
}

func TestChampDelete(t *testing.T) {
	champ := NewChamp()
	champ.Delete(123)
}

func TestChampFind(t *testing.T) {
	champ := NewChamp()
	champ.Find(123)
}

func TestChampFirstRest(t *testing.T) {
	champ := NewChamp()
	champ.FirstRest()
}
