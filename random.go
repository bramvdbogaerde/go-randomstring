package randomstring

import (
	"crypto/rand"
	"math"
	"math/big"
)

// Returns a random integer from 0 up to 2^64-1
func randomInt() int {
	// Currently only works well on 64-bit systems
	max := big.NewInt(int64(math.Pow(2, 64)) - 1)
	r, _ := rand.Int(rand.Reader, max)
	r_int := int(r.Int64())

	return r_int
}
