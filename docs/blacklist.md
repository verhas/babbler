# Blacklist

14 two-syllable words are excluded from encoder output — about 1.08%
of the two-syllable word space (14 of 36 × 36 = 1,296 combinations).
These are filtered out of `GOOD_PAIRS` once, up front (see
[algorithm.md](algorithm.md)), so it's structurally impossible for
`numberToId` to ever produce one — there's no retry logic to reason
about.

Because 14 combinations are removed from *each* syllable position, the
number of valid names shrinks from `36^4 = 1,679,616` to
`1282 × 1282 = 1,643,524` — which is why `MAX_NUM` is `1,643,523`,
not `1,679,615`. See "Why MAX_NUM isn't 36^4 - 1" in
[algorithm.md](algorithm.md).

| Word   | Reason                           |
|--------|----------------------------------|
| `bobo` | fool (French), hurt              |
| `dodo` | brainless bird                   |
| `dudu` | baby poop/pee talk               |
| `gaga` | senile                           |
| `kaka` | poop (German, Dutch, Portuguese) |
| `kuku` | cuckoo                           |
| `lala` | nonsense                         |
| `mumu` | baby poop/pee talk               |
| `popo` | police slang / poop (Spanish)    |
| `soso` | mediocre                         |
| `toto` | toilet slang                     |
| `tutu` | baby poop/pee talk               |
| `nuna` | vulgar for vagina                |
| `suna` | vulgar for vagina                |

## Not blacklisted

`mama`, `papa`, `nana` are intentionally **not** blacklisted — they
aid memorability and carry no offensive connotation in the languages
considered.

## Extending the blacklist

The blacklist is a plain set of lowercase strings, defined once per
language alongside `CV_MAP` (`constants.js`, `constants.py`,
`Constants.java`, `constants.go`). `GOOD_PAIRS`, `GOOD_COUNT`, and
`MAX_NUM` are all *derived* from `BLACKLIST` at load time, so editing
the blacklist automatically and correctly shrinks or grows `MAX_NUM`
— there's no separate value to keep in sync. To add a custom
blacklist for a specific deployment, replace or extend that set — see
the "Custom blacklists per use case" item in the top-level README's
roadmap.
