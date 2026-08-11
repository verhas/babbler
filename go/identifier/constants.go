// Package identifier converts an integer counter into a pronounceable,
// human-friendly name.
package identifier

// CVMap holds the 36 CV syllables, index 0-35. CVMap[0] = "ba", CVMap[35] = "zu".
var CVMap = [36]string{
	"ba", "bo", "bu", "da", "do", "du", "ga", "go", "gu", "ka",
	"ko", "ku", "la", "lo", "lu", "ma", "mo", "mu", "na", "no",
	"nu", "pa", "po", "pu", "ra", "ro", "ru", "sa", "so", "su",
	"ta", "to", "tu", "va", "zo", "zu",
}

// Blacklist holds two-syllable words excluded from encoder output.
var Blacklist = map[string]struct{}{
	"bobo": {}, "dodo": {}, "dudu": {}, "gaga": {}, "kaka": {}, "kuku": {},
	"lala": {}, "mumu": {}, "popo": {}, "soso": {}, "toto": {}, "tutu": {},
	"nuna": {}, "suna": {},
}

// GoodPairs holds all 1,296 two-syllable combinations, filtered down to the
// ones that don't spell a blacklisted word. A name is built from two entries
// of this slice, so it is structurally impossible for NumberToID to produce
// a blacklisted word -- no retry loop needed.
//
// Computed once, in this package-level initializer, from CVMap/Blacklist
// (not hardcoded) so those two stay the single source of truth and this can
// never drift out of sync with them. The cost (under a millisecond) is paid
// once per process, not per call.
var GoodPairs = buildGoodPairs()

// GoodCount is the total number of names NumberToID can produce:
// len(GoodPairs) ^ 2.
var GoodCount = int64(len(GoodPairs)) * int64(len(GoodPairs))

// MaxNum is the highest valid input to NumberToID: GoodCount - 1.
var MaxNum = GoodCount - 1

func buildGoodPairs() []int {
	n := len(CVMap)
	pairs := make([]int, 0, n*n)
	for i := 0; i < n*n; i++ {
		d0, d1 := i/n, i%n
		if _, blacklisted := Blacklist[CVMap[d0]+CVMap[d1]]; !blacklisted {
			pairs = append(pairs, i)
		}
	}
	return pairs
}
