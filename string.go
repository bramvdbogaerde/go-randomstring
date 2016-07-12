package uniquestring

import (
	"github.com/knq/baseconv"
	"strconv"
)

// Returns a random string encoded in Base62
// Its value is generated using the crypto/rand libraries, so it is as secure as your random device/service is
func New() (output string) {
	random_int := randomInt()
	random_intstring := strconv.Itoa(random_int)

	output, _ = baseconv.Convert(random_intstring, baseconv.DigitsDec, baseconv.Digits62)
	return
}
