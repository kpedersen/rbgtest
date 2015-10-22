package rbgtest

import "testing"

func TestRandomBitIsZero(t *testing.T) {
	if RandomBit() == 0 {
		t.Errorf("RandomBit() returned 0")
	}
}
