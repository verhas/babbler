# Identifier Encoder Project — Claude Code Guidelines

## Project Overview

**babbler** is a library for converting integers (typically an
auto-increment counter) into pronounceable, human-friendly
identifiers using consonant-vowel (CV) syllable pairs.

**Core algorithm:**
1. Filter the 1,296 possible two-syllable combinations down to the
   1,282 that don't spell a blacklisted word (`GOOD_PAIRS`) — computed
   once from `CV_MAP`/`BLACKLIST`, not hardcoded.
2. Scramble `num` with a fixed affine permutation:
   `scrambled = (A * num + C) mod GOOD_COUNT`.
3. Split `scrambled` into two independent indices into `GOOD_PAIRS`.
4. Map each to two `CV_MAP` syllables and format as two words:
   `Word1 Word2` (capitalized).

No hashing, no retry loop — steps 2–3 are a genuine mathematical
bijection, so blacklist-avoidance and uniqueness are both guaranteed
by construction, not checked or retried at call time.

**Range:** Numbers 0 to 1,643,523 (`1282^2 - 1` — see "Why MAX_NUM
isn't 36^4 - 1" below for why this isn't `36^4 - 1`)
**Output:** Always exactly 4 syllables (2 per word)
**Example:** 10000 → `Talo Buno`, 10001 → `Patu Luta`

**This is a one-way generator, not a reversible codec.** Given a
counter value it deterministically produces a name; there is no
`idToNumber`/decoder. Unlike earlier versions of this spec, the
forward mapping *is* mathematically invertible (it's a bijection) —
this library simply doesn't expose that direction, because decoding
isn't a requirement (see "Non-goals").

---

## Non-goals

- **No decoding.** There is no `idToNumber(id) → number`. Names are
  display-only labels for a counter/sequence the caller already
  tracks (e.g. a database primary key). Do not add a decoder,
  brute-force search, or reverse lookup back into this library, even
  though the algorithm would technically support one.

---

## Why this algorithm, not "hash and mod"

An earlier version of this spec hashed `num` with MD5 and folded the
result into `[0, 36^4)` with a modulo, retrying with a salted rehash
on a blacklist collision. That has a fatal flaw: hashing `N` numbers
into a same-size bucket space is **not a bijection** — by the
birthday paradox, roughly a third of numbers collided with another
number's name. No retry logic fixes this after the fact, because the
information is genuinely lost the moment two different numbers hash
to the same bucket. This matters because uniqueness (not
reversibility) turned out to be the actual requirement: a caller
bumping a counter and issuing a name needs a guarantee that name has
never been issued before, without checking a database — MD5-and-mod
cannot provide that guarantee no matter how it's tuned.

The current algorithm (affine permutation over the filtered
`GOOD_PAIRS` space) *is* a real bijection, so uniqueness is provable,
not probabilistic. See "Why the affine transform is a bijection"
below, and [docs/algorithm.md](docs/algorithm.md) for the full
derivation.

**Do not reintroduce MD5 or any other hash into the per-call
encoding path.** It adds a dependency and cost for zero benefit here
— the scrambling step only needs `gcd(A, GOOD_COUNT) = 1`, which a
plain hardcoded constant satisfies exactly as well as anything hash-
derived, without the import or the runtime cost.

---

## Syllable Mapping

36 CV pairs (consonant-vowel):
```
ba, bo, bu, da, do, du, ga, go, gu, ka,
ko, ku, la, lo, lu, ma, mo, mu, na, no,
nu, pa, po, pu, ra, ro, ru, sa, so, su,
ta, to, tu, va, zo, zu
```

**Consonants:** b, d, g, k, l, m, n, p, r, s, t, v, z (no h, no confusing pairs)
**Vowels:** a, o, u (distinct, universal across languages)

---

## Blacklist

14 words to exclude from encoding (1.08% of the two-syllable word
space):

```
bobo, dodo, dudu, gaga, kaka, kuku, 
lala, mumu, popo, soso, toto, tutu,
nuna, suna
```

Reasons:
- `bobo` — fool (French), hurt
- `dodo` — brainless bird
- `dudu`, `mumu`, `tutu` — baby poop/pee talk
- `gaga` — senile
- `kaka` — poop (German, Dutch, Portuguese)
- `kuku` — cuckoo
- `lala` — nonsense
- `popo` — police slang / poop (Spanish)
- `soso` — mediocre
- `toto` — toilet slang
- `nuna`, `suna` — vulgar for vagina

**NOT blacklisted (innocent):** `mama`, `papa`, `nana` — these aid memorability and have no offensive connotations.

**Note:** not every blacklisted word is a doubled syllable (`nuna` =
`nu`+`na`, two *different* syllables) — don't assume blacklist
detection can be simplified to "digit0 == digit1".

---

## Implementation Requirements

### Core Function

#### `numberToId(num: number) → string`
- Input: Integer `0`–`MAX_NUM` (`1,643,523`)
- Output: "Word1 Word2" (e.g., "Talo Buno")
- Throws error if num out of range
- Guaranteed unique per `num`, and guaranteed never blacklisted, both
  by construction (see Algorithm Pseudocode) — no retry loop, no
  runtime blacklist check needed

### Constants

**CV_MAP:** Array of 36 syllables in order (index → syllable)

**BLACKLIST:** Set of 14 lowercase problematic words

**GOOD_PAIRS:** All 1,296 two-syllable combinations, filtered down to
the 1,282 that aren't in `BLACKLIST`. **Compute this once, at
module/class load, directly from `CV_MAP` and `BLACKLIST`.** Do not
hardcode it as a separate literal — that would create a second
source of truth that can silently drift out of sync if the blacklist
is ever edited. The computation is ~1,296 iterations (well under a
millisecond); this is cheap enough, and the drift risk bad enough,
that computing it beats hardcoding even for short-lived/ephemeral
processes (CLIs, serverless cold starts).

**GOOD_COUNT:** `len(GOOD_PAIRS)^2` = `1282^2` = `1,643,524`

**MAX_NUM:** `GOOD_COUNT - 1` = `1,643,523`

### Test Coverage

- **Determinism tests:** `numberToId(n)` returns the same string every call for the same `n`
- **Edge cases:** 0, 1, MAX_NUM, out-of-range input (negative and `MAX_NUM + 1`)
- **Format tests:** always exactly two capitalized CV-syllable words
- **Full-range uniqueness + blacklist test:** call `numberToId` for
  every `n` in `[0, MAX_NUM]` (~1.64M calls) and assert zero duplicate
  outputs and zero blacklisted words. This is the test that actually
  proves the core guarantee — sampling is not sufficient. Expect
  roughly 1–2 seconds per language; that's an acceptable, one-time
  test-suite cost, not a per-call cost.

---

## Language-Specific Guidelines

### JavaScript/TypeScript
- Plain `Number` arithmetic is exact here (`A * MAX_NUM` is well
  under `Number.MAX_SAFE_INTEGER`) — no `BigInt` needed
- Export `numberToId`
- Include TypeScript type definitions in `.d.ts` file if TypeScript version

### Python
- Arbitrary-precision ints mean no overflow concerns
- Provide both function and CLI entry point via `__main__`
- Include docstrings (Google style) for all functions

### Java
- Use `long` for the affine transform (`A * num` overflows `int`)
- Create interface `IdEncoder` with default implementation
- Compute `GOOD_PAIRS` in a `Constants` static initializer
- Unit tests with JUnit 5

### Go
- Use `int64` for the affine transform and for `GoodCount`/`MaxNum`
  (portable across 32-bit platforms, where plain `int` may be 32-bit)
- Provide both library and `cmd/id-tool` CLI
- Tests in `*_test.go` files using `testing` package

---

## File Organization

```
{lang}/
├── src/main/encoder.{ext}      # Core encoder
├── src/main/constants.{ext}    # CV map + blacklist + GOOD_PAIRS (single source of truth)
├── test/encoder.test.{ext}     # Encoder tests
└── README.md                    # Language-specific usage
```

---

## Constants (Must be Identical Across All Languages)

```
CV_MAP = [
  'ba', 'bo', 'bu', 'da', 'do', 'du', 'ga', 'go', 'gu', 'ka',
  'ko', 'ku', 'la', 'lo', 'lu', 'ma', 'mo', 'mu', 'na', 'no',
  'nu', 'pa', 'po', 'pu', 'ra', 'ro', 'ru', 'sa', 'so', 'su',
  'ta', 'to', 'tu', 'va', 'zo', 'zu'
]

BLACKLIST = {
  'bobo', 'dodo', 'dudu', 'gaga', 'kaka', 'kuku',
  'lala', 'mumu', 'popo', 'soso', 'toto', 'tutu',
  'nuna', 'suna'
}

# Derived, not hardcoded — see Implementation Requirements above.
GOOD_PAIRS = [i for i in range(36*36)
              if CV_MAP[i // 36] + CV_MAP[i % 36] not in BLACKLIST]   # length 1282
GOOD_COUNT = len(GOOD_PAIRS) ** 2   # 1,643,524
MAX_NUM = GOOD_COUNT - 1            # 1,643,523

# Fixed affine-permutation constants — see "Why the affine transform
# is a bijection" below. Any A coprime to GOOD_COUNT and any C work
# mathematically; these two were picked arbitrarily once and are now
# fixed forever so output is identical across every implementation.
A = 1256797
C = 1443960
```

---

## Algorithm Pseudocode

```
function numberToId(num):
  if num < 0 or num > MAX_NUM:
    throw OutOfRangeError

  scrambled = (A * num + C) mod GOOD_COUNT

  pairCount = len(GOOD_PAIRS)
  outer = scrambled // pairCount
  inner = scrambled % pairCount

  pairA = GOOD_PAIRS[outer]
  pairB = GOOD_PAIRS[inner]

  word1 = CV_MAP[pairA // 36] + CV_MAP[pairA % 36]
  word2 = CV_MAP[pairB // 36] + CV_MAP[pairB % 36]

  return capitalize(word1) + ' ' + capitalize(word2)
```

That's the whole algorithm — no hashing, no loop, no error path other
than the initial range check.

### Why the affine transform is a bijection

`f(x) = (A * x + C) mod n` is a bijection on `Z_n` whenever
`gcd(A, n) = 1` (standard number theory: `A` has a multiplicative
inverse mod `n`, so `f` is invertible). That's the only requirement;
nothing is special about the specific `A`/`C` values beyond satisfying
it. Composed with the `GOOD_PAIRS` split/lookup (also a bijection,
since `outer`/`inner` independently index a fixed-size list), the
whole pipeline is a bijection from `[0, MAX_NUM]` onto the full set of
valid (unique, non-blacklisted) names.

### Why MAX_NUM isn't 36^4 - 1

A name is two picks from `GOOD_PAIRS` (1,282 entries, not all 1,296
combinations — 14 are excluded per position), so there are only
`1282 × 1282 = 1,643,524` possible names. You cannot injectively map
more numbers than there are valid names (pigeonhole principle), so
`MAX_NUM` must be `1,643,524 - 1 = 1,643,523`. If `BLACKLIST` is ever
edited, `MAX_NUM` changes too, automatically, since it's derived —
never hardcode a `MAX_NUM` value independent of `BLACKLIST`.

---

## Testing Checklist

For each implementation, verify:

- [ ] `numberToId(0)` returns valid 4-syllable id
- [ ] `numberToId(1)` returns different id than `numberToId(0)`
- [ ] `numberToId(10000)` returns `"Talo Buno"` and `numberToId(10001)` returns `"Patu Luta"` (reference values for the current `A`/`C`/`BLACKLIST`)
- [ ] `numberToId(MAX_NUM)` returns `"Dobu Zusa"`
- [ ] `numberToId(MAX_NUM + 1)` throws error
- [ ] `numberToId(-1)` throws error
- [ ] Same input always produces the same output (determinism)
- [ ] Output is always exactly two capitalized CV-syllable words
- [ ] **Full range** (`0` to `MAX_NUM`): zero duplicate outputs, zero blacklisted words — this is the test that matters most; see Test Coverage above

---

## Common Pitfalls

1. **Blacklist case sensitivity:** Store lowercase, compare lowercase
2. **Syllable array indexing:** Ensure CV_MAP[0] = 'ba', CV_MAP[35] = 'zu'
3. **Capitalization:** First letter uppercase only (`Talo`, not `TALO` or `talo`)
4. **Integer overflow in the affine transform:** `A * num` can exceed
   32-bit range (e.g. Java `int`) — use a 64-bit type for the
   multiplication and modulo
5. **Don't hardcode `GOOD_PAIRS` or `MAX_NUM`.** Derive both from
   `CV_MAP`/`BLACKLIST` at load time — see Implementation Requirements
6. **Don't try to make this reversible.** No decoder, no reverse index — see "Non-goals" (even though the math would allow it)
7. **Don't reintroduce hashing (MD5 or otherwise) into the per-call path** — see "Why this algorithm, not 'hash and mod'"

---

## Development Workflow

When adding a new language:

1. Copy constants from existing implementation (don't redefine) —
   `CV_MAP`, `BLACKLIST`, `A`, `C` must be byte-for-byte identical;
   `GOOD_PAIRS`/`GOOD_COUNT`/`MAX_NUM` must be *derived*, not copied
2. Implement `numberToId()`
3. Run tests against the same test suite used in other languages
4. Verify determinism and the full-range uniqueness/blacklist test
5. Document any language-specific notes in the language's README.md
6. Update top-level README.md with language support status

---

## Performance Notes

- **Encoding:** O(1) — one multiplication, one division, one modulo, two array lookups. No hashing at call time.
- **GOOD_PAIRS setup:** ~1,296 iterations, paid once per process at load time (measured ~0.4ms cold in JavaScript) — see [docs/performance.md](docs/performance.md)

---

## Future Enhancements

- [ ] Web UI for generating names
- [ ] REST API endpoint
- [ ] Batch generation tool
- [ ] Performance benchmarks across languages
- [ ] Support for custom blacklists per use case (remember: changing
      `BLACKLIST` changes `MAX_NUM` too, since it's derived)

---

## Contact / Repository

**GitHub:** `verhas/babbler`
**Author:** Peter Verhas (@verhas)
**License:** MIT (or your choice)

---

## Revision History

- v1.2: Replaced MD5-hash-then-retry with a provable bijection: a
  fixed affine permutation (`(A*num+C) mod GOOD_COUNT`) over
  `GOOD_PAIRS`, the blacklist-filtered set of two-syllable
  combinations. This guarantees both uniqueness (no two numbers ever
  produce the same name) and blacklist-avoidance by construction, with
  no retry loop and no hashing at call time. `MAX_NUM` changed from
  `1,679,615` to `1,643,523` — the exact count of valid names is
  smaller than `36^4` because blacklisted combinations are excluded
  from both syllable positions (pigeonhole principle: you can't
  injectively map more numbers than there are valid names). This was
  necessary because the actual requirement was uniqueness without a
  database check ("bump a counter, get a new unique name, without
  checking if the given ID was already issued") — MD5-and-mod is
  fundamentally incapable of that guarantee, and salting the retry
  (v1.1) only fixed blacklist-avoidance, not uniqueness.
- v1.1: Dropped `idToNumber`/decoder — reversibility is not a
  requirement; the library is a one-way name generator over a caller-
  supplied counter. Fixed the blacklist-retry algorithm to salt the
  MD5 input per attempt instead of incrementing the scrambled integer
  (superseded by v1.2, which removes the retry loop entirely).
- v1.0 (initial): JavaScript, Python, Java implementations with MD5-based scrambling and a brute-force decoder (since removed, see v1.1).
