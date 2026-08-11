package identifier

import (
	"fmt"
	"strings"
)

// Fixed affine-permutation constants. A modular multiplication
// x -> A*x + C (mod GoodCount) is a bijection on Z_GoodCount whenever
// gcd(A, GoodCount) is 1 -- that's the only requirement. These two were
// picked arbitrarily once and are now fixed forever so output is stable
// across runs and identical across every language implementation.
const permA int64 = 1256797
const permC int64 = 1443960

func pairToWord(pairIndex int) string {
	n := len(CVMap)
	return CVMap[pairIndex/n] + CVMap[pairIndex%n]
}

func capitalize(word string) string {
	return strings.ToUpper(word[:1]) + word[1:]
}

// NumberToID converts an integer counter into a pronounceable, human-friendly
// identifier such as "Talo Buno". num must be in [0, MaxNum].
//
// Guaranteed unique per num and never a blacklisted word, both by
// construction (see GoodPairs), not by retrying. This is a one-way
// generator: there is no decoder (see the top-level README's "Non-goals").
func NumberToID(num int) (string, error) {
	if num < 0 || int64(num) > MaxNum {
		return "", fmt.Errorf("num must be in [0, %d], got %d", MaxNum, num)
	}

	pairCount := int64(len(GoodPairs))
	scrambled := (permA*int64(num) + permC) % GoodCount
	outer := int(scrambled / pairCount)
	inner := int(scrambled % pairCount)

	word1 := pairToWord(GoodPairs[outer])
	word2 := pairToWord(GoodPairs[inner])

	return capitalize(word1) + " " + capitalize(word2), nil
}
