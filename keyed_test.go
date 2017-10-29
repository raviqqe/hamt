package champ

import "testing"

type KeyedInt int32

func (i KeyedInt) Key() int32 {
	return int32(i)
}

func TestKeyed(t *testing.T) {
	h := Keyed(KeyedInt(42))
	t.Log(h)
}
