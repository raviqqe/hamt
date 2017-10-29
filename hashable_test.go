package champ

import "testing"

type HashableInt int32

func (i HashableInt) Hash() int32 {
	return int32(i)
}

func TestHashable(t *testing.T) {
	h := Hashable(HashableInt(42))
	t.Log(h)
}
