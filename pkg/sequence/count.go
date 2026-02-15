package sequence

import "iter"

// Numeric is a constraint for types that support addition.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Count returns an infinite iterator that produces evenly spaced values
// beginning with start and incrementing by step.
//
// Example usage:
//
//	for i, val := range sequence.Count(10, 1) {
//	    if i >= 5 { break }
//	    fmt.Println(val) // 10, 11, 12, 13, 14
//	}
//
//	for i, val := range sequence.Count(2.5, 0.5) {
//	    if i >= 3 { break }
//	    fmt.Println(val) // 2.5, 3.0, 3.5
//	}
func Count[N Numeric](start, step N) iter.Seq[N] {
	return func(yield func(N) bool) {
		n := start
		for {
			if !yield(n) {
				return
			}
			n += step
		}
	}
}
