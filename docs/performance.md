# Performance

## Encoding — O(1), no hashing at call time

`numberToId` is one modular multiplication, one division, one modulo,
and two array lookups — no hashing, no retry loop. Cost is constant
regardless of the input number.

## One-time setup cost: building GOOD_PAIRS

`GOOD_PAIRS` (the 1,282 non-blacklisted syllable pairs) is computed
once, at module/class load, by filtering all 1,296 combinations
against `BLACKLIST`. Measured in JavaScript: about 0.4ms on a single
cold call, ~0.08ms once the JIT warms up — negligible next to typical
process/runtime startup (tens to hundreds of ms), and it's paid once
per process, not once per call. It's deliberately *not* precomputed
into a hardcoded literal: the computation is cheap enough that doing
it at load time removes any risk of the literal drifting out of sync
with `CV_MAP`/`BLACKLIST` if either is ever edited, for near-zero
practical cost — see [algorithm.md](algorithm.md).

## No decode cost to worry about

`babbler` is a one-way generator (see the top-level README's
"Non-goals"): there is no `idToNumber`, so there's no brute-force
search or reverse-index to budget for. If your application needs to
look a name back up to its number, store the pair yourself (e.g. as
two columns in a database row).

## Benchmark checklist

Each language's test suite calls `numberToId` for every number in
`[0, MAX_NUM]` (about 1.64M calls) to directly verify zero collisions
and zero blacklisted words — this doubles as a performance sanity
check. Measured full-range runtimes: JavaScript ~1.4s, Python ~1.7s,
Java ~0.9s, Go ~0.8s.
