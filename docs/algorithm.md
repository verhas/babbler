# Algorithm

`babbler` turns an integer counter into a pair of pronounceable,
capitalized CV (consonant-vowel) words. The same steps are
implemented identically in every language in this repo.

**Two guarantees, both structural (not checked at runtime):**

- **Uniqueness.** Every `num` in `[0, MAX_NUM]` produces a distinct
  name — no two different numbers ever collide.
- **No blacklisted words.** No name is ever one of the 14 blacklisted
  words (see [blacklist.md](blacklist.md)).

## Why this needs more than "hash and mod"

An earlier version of this algorithm hashed `num` with MD5 and folded
the result into `[0, 36^4)` with a modulo. That is **not** a
bijection: hashing `N` numbers into a same-size bucket space
necessarily collides for a large fraction of them (the birthday
paradox — empirically, about 37% of numbers collided with another
number's name). No amount of retrying fixes this after the fact; the
information is lost at the moment two different numbers hash to the
same bucket.

The current algorithm avoids this by construction instead: it builds
a real bijection over exactly the set of valid (non-blacklisted)
names, so collisions are mathematically impossible rather than merely
unlikely.

## Constants

- **CV_MAP** — 36 two-letter syllables (`ba`, `bo`, `bu`, ... `zo`,
  `zu`), index 0 to 35.
- **BLACKLIST** — 14 two-syllable words excluded from output (see
  [blacklist.md](blacklist.md)).
- **GOOD_PAIRS** — all 1,296 two-syllable combinations (`36 × 36`),
  filtered down to the 1,282 that don't spell a blacklisted word.
  Computed once, at module/class load, directly from `CV_MAP` and
  `BLACKLIST` (never hardcoded), so it can't drift out of sync with
  them. The cost of computing it (under a millisecond) is paid once
  per process, not per call — see the note in each language's
  `constants` file.
- **GOOD_COUNT** — `len(GOOD_PAIRS)^2 = 1282^2 = 1,643,524`. This is
  the total number of names the algorithm can produce, since a name
  is just two independent picks from `GOOD_PAIRS`.
- **MAX_NUM** — `GOOD_COUNT - 1 = 1,643,523`. Valid input range is
  `[0, MAX_NUM]`. This is *smaller* than the old `36^4 - 1`
  (`1,679,615`) by exactly the number of two-syllable slots the
  blacklist removes (14 of 1,296) — see "Why MAX_NUM isn't 36^4 - 1"
  below.

## Encoding: `numberToId(num)`

1. **Validate range.** If `num < 0` or `num > MAX_NUM`, raise an
   out-of-range error.
2. **Scramble.** `scrambled = (A * num + C) mod GOOD_COUNT`, a fixed
   affine transform. This is the only "randomizing" step, and it
   replaces MD5 entirely — see "Why the affine transform is a
   bijection" below.
3. **Split.** `outer = scrambled // len(GOOD_PAIRS)`,
   `inner = scrambled % len(GOOD_PAIRS)`.
4. **Look up two pairs.** `pairA = GOOD_PAIRS[outer]`,
   `pairB = GOOD_PAIRS[inner]`. Each is an index in `[0, 1296)`;
   convert to two `CV_MAP` digits: `d0 = pairA // 36`,
   `d1 = pairA % 36` (similarly `d2`, `d3` for `pairB`).
5. **Map to syllables and form words.**
   `word1 = CV_MAP[d0] + CV_MAP[d1]`,
   `word2 = CV_MAP[d2] + CV_MAP[d3]`. Neither can be blacklisted —
   `GOOD_PAIRS` only contains indices that were already filtered to
   exclude blacklisted syllable pairs.
6. **Capitalize.** Return `Capitalize(word1) + " " + Capitalize(word2)`
   (first letter uppercase only, e.g. `Talo`, not `TALO`).

No retry loop, no hashing at call time: this is pure integer
arithmetic plus two array lookups.

### Why the affine transform is a bijection

`f(x) = (A * x + C) mod n` is a bijection on `Z_n` (the integers
`0` to `n-1`) whenever `gcd(A, n) = 1` — a standard, easily-checked
number-theory fact (`A` has a multiplicative inverse mod `n`, so `f`
is invertible). That's the *only* requirement on `A`; there is
nothing special about the specific values chosen. `A = 1256797` and
`C = 1443960` were picked once, arbitrarily (verified to satisfy
`gcd(A, GOOD_COUNT) = 1`), and are now fixed constants hardcoded
identically in every language — the same `num` must map to the same
`scrambled` value everywhere.

Because step 2 is a bijection on `[0, GOOD_COUNT)`, and steps 3–6 are
a bijection between `[0, GOOD_COUNT)` and the set of valid names
(splitting into `(outer, inner)` and independently indexing
`GOOD_PAIRS` for each), the composition is a bijection from
`[0, MAX_NUM]` onto the full set of valid names. That's what
guarantees uniqueness — it's provable, and also checked directly: each
language's test suite calls `numberToId` for **every** number in
`[0, MAX_NUM]` and asserts zero duplicate names and zero blacklisted
words (takes about 1–1.5 seconds per language).

### Why MAX_NUM isn't 36^4 - 1

A name is two picks from `GOOD_PAIRS` (1,282 non-blacklisted
syllable-pairs), so there are only `1282 × 1282 = 1,643,524` possible
names — fewer than the `36^4 = 1,679,616` possible *raw* 4-digit
combinations, because blacklisted combinations are excluded from both
halves. You cannot injectively map more numbers than there are valid
names (pigeonhole principle), so `MAX_NUM` must be
`1,643,524 - 1 = 1,643,523`, not `36^4 - 1`.

## Worked example

`numberToId(10000)`:

1. `scrambled = (1256797 * 10000 + 1443960) mod 1643524 = 1385932`
2. `outer = 1385932 // 1282 = 1081`, `inner = 1385932 % 1282 = 90`
3. `GOOD_PAIRS[1081]` and `GOOD_PAIRS[90]` look up to two digit pairs,
   which map through `CV_MAP` to `ta`+`lo` and `bu`+`no`.
4. Words: `talo`, `buno` — neither is blacklisted (guaranteed).
5. Result: `"Talo Buno"`.
