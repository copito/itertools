# itertools

A Go library inspired by Python's itertools, providing functional iteration utilities that leverage Go's native `iter` package (Go 1.23+).

## Overview

This library brings the elegance and power of Python's itertools to Go, offering a simplified and idiomatic way for Go developers to work with iterators. It extends Go's `iter` package with additional functionality for grouping, combinatorics, and infinite sequences.

## Features

### Currently Implemented

#### GroupBy
Groups consecutive elements by a key function, mirroring Python's `itertools.groupby()`.

```go
import (
    "fmt"
    "slices"
    "github.com/copito/itertools/pkg/group"
)

data := []string{"apple", "apricot", "banana", "berry", "cherry"}

for key, group := range group.GroupBy(slices.Values(data), func(s string) rune {
    return rune(s[0]) // group by first letter
}) {
    fmt.Printf("%c: ", key)
    for val := range group {
        fmt.Printf("%s ", val)
    }
    fmt.Println()
}
// Output:
// a: apple apricot
// b: banana berry
// c: cherry
```

**Important**: Like Python's `groupby()`, this function only groups *consecutive* elements with the same key. Pre-sort your data if you want all elements with the same key grouped together.

### Planned Features

#### Combinatoric Iterators
- `Combinations` - r-length subsequences of elements from the input
- `Permutations` - r-length permutations of elements
- `Product` - Cartesian product of input iterables

#### Infinite Iterators
- `Count` - infinite counter starting from a value
- `Cycle` - infinite repetition of an iterable
- `Repeat` - repeat an element indefinitely or n times

## Installation

```bash
go get github.com/copito/itertools
```

**Requirements**: Go 1.23 or higher (for `iter` package support)

## Usage Examples

### Working with Structs

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Charlie", 30},
    {"Bob", 25},
    {"David", 25},
}

// Group by age
for age, group := range group.GroupBy(slices.Values(people), func(p Person) int {
    return p.Age
}) {
    fmt.Printf("Age %d: ", age)
    for person := range group {
        fmt.Printf("%s ", person.Name)
    }
    fmt.Println()
}
// Output:
// Age 30: Alice Charlie
// Age 25: Bob David
```

### Complex Keys

```go
type AgeGroup struct {
    Decade      int
    LastLetter  rune
}

for key, group := range group.GroupBy(slices.Values(people), func(p Person) AgeGroup {
    return AgeGroup{
        Decade:     p.Age / 10 * 10,
        LastLetter: rune(p.Name[len(p.Name)-1]),
    }
}) {
    fmt.Printf("Group %+v: ", key)
    for person := range group {
        fmt.Printf("%s ", person.Name)
    }
    fmt.Println()
}
```

## Design Philosophy

This library aims to:

1. **Stay idiomatic to Go** - Use Go's native `iter.Seq` and `iter.Seq2` types
2. **Maintain Python familiarity** - Keep similar semantics and behavior where appropriate
3. **Leverage generics** - Provide type-safe operations for any data type
4. **Be memory efficient** - Use lazy evaluation through iterators
5. **Remain simple** - Provide clean, easy-to-understand APIs

## Comparison with Python's itertools

| Python | Go (this library) | Status |
|--------|------------------|--------|
| `itertools.groupby()` | `group.GroupBy()` | ✅ Implemented |
| `itertools.combinations()` | `combinatoric.Combinations()` | 🚧 Planned |
| `itertools.permutations()` | `combinatoric.Permutations()` | 🚧 Planned |
| `itertools.product()` | `combinatoric.Product()` | 🚧 Planned |
| `itertools.count()` | `sequence.Count()` | 🚧 Planned |
| `itertools.cycle()` | `sequence.Cycle()` | 🚧 Planned |
| `itertools.repeat()` | `sequence.Repeat()` | 🚧 Planned |

## Contributing

Contributions are welcome! This project is in active development. Feel free to:

- Report bugs
- Suggest new features
- Submit pull requests
- Improve documentation

## License

See [LICENSE](LICENSE) file for details.

## Why itertools for Go?

Go's `iter` package (introduced in Go 1.23) provides the foundation for functional iteration patterns, but it lacks the rich set of utilities that Python developers have come to love. This library fills that gap, making it easier to:

- Process data streams efficiently
- Work with combinatorial problems
- Handle infinite sequences
- Write more functional and composable code

Whether you're coming from Python or just want more powerful iteration tools in Go, this library provides a familiar and efficient way to work with sequences.
