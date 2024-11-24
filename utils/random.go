package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rng *rand.Rand

// Initialize the random number generator once with a random seed
func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(source)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString() string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < len(alphabet); i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rng.Intn(len(currencies))]
}
