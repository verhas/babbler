# babbler (Go)

Convert a counter into a pronounceable, human-friendly name.

```bash
go test ./...
go run ./cmd/id-tool encode 10000
```

Add as a dependency:

```bash
go get github.com/verhas/babbler/identifier@latest
```

```go
import "github.com/verhas/babbler/identifier"

identifier.NumberToID(10000) // "Talo Buno", nil
identifier.NumberToID(10001) // "Patu Luta", nil
```

Published on [pkg.go.dev](https://pkg.go.dev/github.com/verhas/babbler/identifier)
— see [release.sh](release.sh) for how new versions get published (a
git tag, not an upload).

## API

### `identifier.NumberToID(num int) (string, error)`

- `num` must be an integer in `[0, 1643523]` (`identifier.MaxNum`),
  or an error is returned.
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

`identifier.CVMap`, `identifier.Blacklist`, and `identifier.MaxNum`
are also exported for consumers that want to reason about the output
space directly.
