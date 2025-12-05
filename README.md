# GoDash

GoDash is a small collection of utility functions for working with slices and maps in Go —
inspired by functional helpers like Lodash. The library provides two flavors of many helpers:

- Original: full-context iterators that receive the element, index (where applicable), and the
	entire collection (slice or map). These are useful when the callback needs context such as
	the current index or the whole collection.
- Simplified: lightweight iterators that only receive the element value and are more convenient
	when you only need the element itself.

## Installation

Install with the standard Go tooling:

```bash
go get github.com/duhnnie/godash@latest
```

## Quick usage examples

Map (original, full context):

```go
import "github.com/duhnnie/godash"

persons := []person{{"Kurt","Cobain"}, {"Dave","Grohl"}}
res, err := godash.Map(persons, func(p person, idx int, coll []person) (string, error) {
		return fmt.Sprintf("%s %s", p.FirstName, p.LastName), nil
})
```

MapTo (simplified):

```go
res, err := godash.MapTo(persons, func(p person) (string, error) {
		return fmt.Sprintf("%s %s", p.FirstName, p.LastName), nil
})
```

## Notes on error handling

Most helpers return an `error` as the second return value. For functions that build
collections (e.g. `Map` / `MapTo`), when a callback returns an error the helper
stops and returns the partially populated result along with the error.

## Original vs Simplified: what differs

- Original functions pass more context to the callback: element, index (for slices), and
	the collection (slice or map). This allows callbacks to inspect neighbors, length or keys.
- Simplified functions only pass the element; callbacks are therefore easier to write when
	they only depend on the element value.
- Both flavors return errors when callbacks return errors; behaviour on error is to stop
	iteration and return any accumulated/partial result together with the error.

## Function mapping (Original → Simplified)

| Original (full context) | Simplified (element-only) |
|---|---|
| `Map` | `MapTo` |
| `Every` | `Everyone` |
| `Some` | `Any` |
| `Find` | `FindFirst` |
| `FindAll` | `Filter` |
| `Reduce` | `ReduceTo` |
| `ReduceMap` | `ReduceMapTo` |

## Other helpers

- `Max`, `Min`, `Clamp` — numeric utilities for common comparisons and clamping values.

## Contributing

Contributions are welcome. Run the tests with:

```bash
go test ./...
```

If you add parallel test execution, use `go test -race ./...` to check for races.

## License

[MIT](./LICENSE)