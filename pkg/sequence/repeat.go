package sequence

import "iter"

// Repeat returns an iterator that yields the same object over and over again.
// If times is provided (as a single variadic argument), it repeats the object
// that many times. Otherwise, it repeats indefinitely.
//
// Example usage:
//
//	// Infinite repeat
//	for i, val := range sequence.Repeat(10) {
//	    if i >= 3 { break }
//	    fmt.Println(val) // 10 10 10
//	}
//
//	// Limited repeat
//	for val := range sequence.Repeat(10, 3) {
//	    fmt.Println(val) // 10 10 10
//	}
//
// Common use case with map or zip operations:
//
//	// Pair each element with a constant value
//	for num, constant := range zip(slices.Values([]int{1,2,3}), sequence.Repeat(5, 3)) {
//	    fmt.Println(num, constant) // (1,5) (2,5) (3,5)
//	}
func Repeat[V any](object V, times ...int) iter.Seq[V] {
	return func(yield func(V) bool) {
		// If times is not specified, repeat indefinitely
		if len(times) == 0 {
			for {
				if !yield(object) {
					return
				}
			}
		} else {
			// Repeat for the specified number of times
			count := times[0]
			for i := 0; i < count; i++ {
				if !yield(object) {
					return
				}
			}
		}
	}
}
