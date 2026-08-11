# Changelog

## v1.2

- Replaced MD5-hash-then-retry with a provable bijection: a fixed
  affine permutation (`(A*num+C) mod GOOD_COUNT`) over `GOOD_PAIRS`,
  the blacklist-filtered set of two-syllable combinations. Every
  number in the valid range now produces a **guaranteed-unique**,
  guaranteed-non-blacklisted name — no retry loop, no hashing at call
  time, and no database check needed to confirm a name hasn't already
  been issued.
- `MAX_NUM` changed from `1,679,615` to `1,643,523`: only
  `1282 × 1282 = 1,643,524` names are possible once blacklisted
  syllable pairs are excluded from both positions, so the usable range
  had to shrink to match exactly (pigeonhole principle) — see
  `docs/algorithm.md`.
- Dropped the MD5/crypto dependency entirely; `numberToId` is now pure
  integer arithmetic in all four languages.
- Each language's test suite now calls `numberToId` for every number
  in the valid range (~1.64M calls, ~1–2s) and asserts zero
  duplicates and zero blacklisted words, directly proving the
  uniqueness guarantee rather than sampling for it.

## v1.1

- Dropped `idToNumber`/decoder — reversibility isn't a requirement,
  and hashing into a same-size output space isn't a bijection (~37%
  of numbers would collide with another number's name). `babbler` is
  now a one-way name generator with no uniqueness guarantee across
  numbers. See the top-level README's "Non-goals".
- Fixed the blacklist retry to salt the MD5 input per attempt instead
  of incrementing the scrambled integer, which could get permanently
  stuck when the *leading* syllable pair collided with the blacklist
  (e.g. `numberToId(145)`).

## v1.0 (initial)

- JavaScript/TypeScript, Python, Java, and Go implementations of the
  MD5-based number ↔ pronounceable-ID encoding.
- Shared constants (`CV_MAP`, blacklist, `MAX_NUM`) across all
  languages.
- Round-trip, blacklist, and edge-case test suites per language.
- Docs: algorithm, usage, performance, blacklist rationale.
- CI running all four language test suites on every push.
