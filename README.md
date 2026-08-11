# babbler

Convert integers — typically an auto-increment counter — into
pronounceable, human-friendly names using consonant-vowel (CV)
syllable pairs.

```
10000 → "Talo Buno"
10001 → "Patu Luta"
```

Instead of showing users `10000`, show them `Talo Buno`: easier to
read aloud, type, and remember. Every number in the valid range gets
a **guaranteed-unique** name — no two different numbers ever collide,
and no name is ever an awkward word — both by construction, not by
checking a database.

## How it works

1. Take the 1,282 two-syllable combinations (of 1,296 possible) that
   don't spell a blacklisted word (`bobo`, `dodo`, `kaka`, ...) —
   computed once from the syllable map and blacklist, not hardcoded.
2. Scramble the input number with a fixed modular-arithmetic
   permutation (`(A * num + C) mod GOOD_COUNT`) — a genuine bijection,
   so distinct numbers always land on distinct outputs.
3. Split the result into two independent picks from that filtered
   list, and format them as two capitalized words, e.g. `Talo Buno`.

No hashing, no retry loop, no reversibility, no database lookups —
just arithmetic. See [docs/algorithm.md](docs/algorithm.md) for
exactly why this guarantees uniqueness and blacklist-avoidance, with a
proof and a full-range test to back it up.

Full details: [docs/algorithm.md](docs/algorithm.md) ·
[docs/blacklist.md](docs/blacklist.md) ·
[docs/performance.md](docs/performance.md) ·
[docs/usage.md](docs/usage.md)

- **Range:** 0 to 1,643,523 (see [docs/algorithm.md](docs/algorithm.md)
  for why this isn't `36^4 - 1`)
- **Output:** always exactly 4 syllables (2 per word)
- **Encoding:** O(1), no hashing at call time

## Non-goals

- **No decoding.** There is no `idToNumber`. Names are display-only
  labels for a counter/sequence you already track (e.g. a database
  primary key) — not a reversible codec. (The forward mapping *is*
  mathematically invertible, since it's a bijection — `babbler` just
  doesn't expose that direction, because nothing needs it.)

## Language support

| Language              | Status | Path          |
|-----------------------|--------|---------------|
| JavaScript/TypeScript | ✅     | [`javascript/`](javascript/) |
| Python                | ✅     | [`python/`](python/)         |
| Java                  | ✅     | [`java/`](java/)             |
| Go                    | ✅     | [`go/`](go/)                 |

All four implementations share the same constants (`CV_MAP`,
blacklist, `MAX_NUM`) and the same algorithm, so the same input number
produces the identical name in every language. Each language's test
suite calls `numberToId` for every number in the valid range (~1.64M
calls) and asserts zero duplicate names and zero blacklisted words.
See [docs/usage.md](docs/usage.md) for per-language install/run
instructions, and [`examples/`](examples/) for minimal runnable
demos.

## Development

See each language's `README.md` for its own build/test commands, and
[CLAUDE.md](CLAUDE.md) for the full project specification followed by
every implementation.

## License

MIT — see [LICENSE](LICENSE).
