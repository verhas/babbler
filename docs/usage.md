# Usage

## JavaScript / TypeScript

```bash
cd javascript
npm test
```

```js
const { numberToId } = require('./src');

numberToId(10000); // "Talo Buno"
numberToId(10001); // "Patu Luta"
```

## Python

```bash
cd python
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
```

Published on PyPI as `javax0-babbler` (the plain name `babbler` was
already taken), but the importable package is `babbler`.

## Java

```bash
cd java
mvn test
```

```java
import com.identifier.encoder.Encoder;
import com.identifier.encoder.IdEncoder;

IdEncoder encoder = new Encoder();
encoder.numberToId(10000); // "Talo Buno"
encoder.numberToId(10001); // "Patu Luta"
```

## Go

```bash
cd go
go test ./...
go run ./cmd/id-tool encode 10000
```

```bash
go get github.com/verhas/babbler/identifier@latest
```

```go
import "github.com/verhas/babbler/identifier"

identifier.NumberToID(10000) // "Talo Buno"
identifier.NumberToID(10001) // "Patu Luta"
```

## Contract across languages

Every implementation exposes the same `numberToId` operation and the
same constants (`CV_MAP`, `BLACKLIST`, `MAX_NUM = 1,643,523`), so the
same input number produces the identical name in every language, and
every number in the valid range produces a guaranteed-unique,
guaranteed-non-blacklisted name. See [algorithm.md](algorithm.md) for
the exact steps and why those guarantees hold.

This is a one-way generator — there is no decoder. See the top-level
[README](../README.md#non-goals) for why.
