# babbler (Java)

Convert a counter into a pronounceable, human-friendly name.

```bash
mvn test
```

```java
import com.javax0.babbler.Encoder;
import com.javax0.babbler.IdEncoder;

IdEncoder encoder = new Encoder();
encoder.numberToId(10000); // "Talo Buno"
encoder.numberToId(10001); // "Patu Luta"
```

## API

### `IdEncoder.numberToId(int num): String`

- `num` must be an integer in `[0, 1643523]` (`Constants.MAX_NUM`), or
  an `IllegalArgumentException` is thrown.
- Deterministic: the same `num` always returns the same string.
- **Unique:** every `num` in the valid range produces a distinct
  name — no two different numbers ever collide, and no name is ever
  a blacklisted word. Both are structural guarantees (see
  [../docs/algorithm.md](../docs/algorithm.md)), not something the
  caller needs to check.
- One-way: there is no decoder. Names are display-only labels for a
  counter/sequence you already track — see the top-level
  [README](../README.md).

`Encoder` implements the `IdEncoder` interface, which carries the
actual algorithm as a default method — implement `IdEncoder` directly
if you want a custom encoder variant.

## Constants

`Constants.CV_MAP`, `Constants.BLACKLIST`, and `Constants.MAX_NUM` are
public for consumers that want to reason about the output space
directly.
