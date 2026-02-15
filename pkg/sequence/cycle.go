package sequence

import "iter"

// Cycle returns an infinite iterator that repeatedly cycles through
// the elements of the input sequence. It saves a copy of each element
// during the first pass and then repeats the saved elements indefinitely.
//
// Example usage:
//
//	data := []string{"A", "B", "C", "D"}
//	for i, val := range sequence.Cycle(slices.Values(data)) {
//	    if i >= 10 { break }
//	    fmt.Println(val) // A B C D A B C D A B
//	}
func Cycle[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		var saved []V

		// First pass: yield elements and save them
		for element := range seq {
			if !yield(element) {
				return
			}
			saved = append(saved, element)
		}

		// If the input was empty, return immediately
		if len(saved) == 0 {
			return
		}

		// Infinite loop: keep cycling through saved elements
		for {
			for _, element := range saved {
				if !yield(element) {
					return
				}
			}
		}
	}
}
