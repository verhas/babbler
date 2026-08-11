# babbler (Python)

Convert a counter into a pronounceable, human-friendly name.

Published on PyPI as `javax0-babbler` (the plain name `babbler` was
already taken) — the distribution name is namespaced, but the
importable package is just `babbler`, matching the `com.javax0`
groupId used for the Java build.

```bash
pip install -e ".[dev]"
pytest
```

```python
from babbler import number_to_id

number_to_id(10000)  # "Talo Buno"
number_to_id(10001)  # "Patu Luta"
```

CLI:

```bash
python -m babbler encode 10000
# or, once installed: babbler encode 10000
```

## API

### `number_to_id(num: int) -> str`

- `num` must be an integer in `[0, 1643523]` (`MAX_NUM`), or a
  `ValueError` is raised.
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
