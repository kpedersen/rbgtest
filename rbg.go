package rbgtest

import (
	"math/rand"
	"time"
)

// RandomBit returns 0 or 1
func RandomBit() uint8 {
	rbg := rand.New(rand.NewSource(time.Now().Unix()))
	return uint8(rbg.Intn(2))
}
