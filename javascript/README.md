# babbler (JavaScript / TypeScript)

Convert a counter into a pronounceable, human-friendly name.

```bash
npm install
npm test
```

```js
const { numberToId } = require('babbler');

numberToId(10000); // "Talo Buno"
numberToId(10001); // "Patu Luta"
```

TypeScript types are provided via `src/index.d.ts`.

## API

### `numberToId(num: number): string`

- `num` must be an integer in `[0, 1643523]` (`MAX_NUM`), or a
  `RangeError` is thrown.
- Deterministic: the same `num` always returns the same string.
- **Unique:** every `num` in the valid range produces a distinct
  name — no two different numbers ever collide, and no name is ever
  a blacklisted word. Both are structural guarantees (see
  [../docs/algorithm.md](../docs/algorithm.md)), not something the
  caller needs to check.
- One-way: there is no decoder. Names are display-only labels for a
  counter/sequence you already track — see the top-level
  [README](../README.md).

## Constants

`CV_MAP`, `BLACKLIST`, and `MAX_NUM` are also exported for consumers
that want to reason about the output space directly.
