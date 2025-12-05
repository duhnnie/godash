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

## No-error (NE) variants

For convenience, the library provides `NE` (no-error) variants for many helpers. These
functions follow the same naming as the original helpers but with an `NE` suffix (for example
`MapNE`, `MapToNE`, `AnyNE`, `EveryoneNE`, `FindNE`, `FindFirstNE`, `FilterNE`, `ReduceNE`,
`ReduceToNE`, `ReduceMapNE`, `ReduceMapToNE`).

### What they do:

- Accept iterator/reducer callbacks that do not return an `error` (callback signatures are simpler).
- Return only the primary result (no `error` return value).

### When to use NE variants:

- Use NE helpers when your callback cannot fail (pure transformation/predicate) and you prefer a
	simpler API without dealing with `error` values.
- Avoid NE helpers when your callback may fail, needs to propagate errors, or when you want to
	stop iteration early with an error. In those cases use the original functions which return `(result, error)`.

### Examples

MapNE (no-error mapper):

```go
persons := []person{{"Kurt","Cobain"}, {"Dave","Grohl"}}
names := godash.MapNE(persons, func(p person, idx int, coll []person) string {
		return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
})
// names is []string{"Kurt Cobain", "Dave Grohl"}
```

MapToNE (element-only no-error mapper):

```go
names := godash.MapToNE(persons, func(p person) string {
		return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
})
```

AnyNE / EveryoneNE:

```go
hasEven := godash.AnyNE([]int{1, 2, 3}, func(n int) bool { return n%2 == 0 })
allEven := godash.EveryoneNE([]int{2, 4, 6}, func(n int) bool { return n%2 == 0 })
```

ReduceNE / ReduceToNE:

```go
sum := godash.ReduceNE([]int{1, 2, 3}, func(acc, cur, _ int, _ []int) int { return acc + cur }, 0)
// or element-only
sum2 := godash.ReduceToNE([]int{1, 2, 3}, func(acc int, cur int) int { return acc + cur }, 0)
```

FindNE / FindFirstNE:

```go
item, found := godash.FindNE([]int{1, 2, 3}, func(it int, idx int, s []int) bool { return it == 2 })
if found { /* use item */ }
```

Note: NE variants intentionally do not provide error propagation. If you need errors (for
example validation or IO during mapping), use the original versions so you can return and
handle errors properly.

## Contributing

Contributions are welcome. Run the tests with:

```bash
go test ./...
```

If you add parallel test execution, use `go test -race ./...` to check for races.

## License

[MIT](./LICENSE)