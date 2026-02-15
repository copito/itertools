package group

import "iter"

// Group By is a function that groups elements of a slice based on a key function.
// Importantly, it only groups *consecutive* elements with the same key, so the input
// should be pre-sorted by the key function if you want all elements with the same key to be grouped together.
//
// Example usage:
//
//	for key, group := range GroupBy(slices.Values(s), keyFunc) {
//	    for val := range group {
//	        fmt.Println(key, val)
//	    }
//	}
func GroupBy[V any, K comparable](seq iter.Seq[V], keyFunc func(V) K) iter.Seq2[K, iter.Seq[V]] {
	return func(yield func(K, iter.Seq[V]) bool) {
		// Shared mutable state between the outer loop and the inner grouper,
		// analogous to Python's `nonlocal curr_value, curr_key, exhausted`.
		var (
			currValue V
			currKey   K
			exhausted = false
		)

		// Pull a pull-style iterator from then push-style iter.Seq.
		// `next()` returns (value, ok) - ok is false when the input is exhausted.
		next, stop := iter.Pull(seq)
		defer stop()

		// Try to get the first element. If the input is empty, return immediately.
		v, ok := next()
		if !ok {
			return
		}
		currValue = v
		currKey = keyFunc(currValue)

		// _grouper yields all consecutive elements whose key equals targetKey.
		// It updates currValue/currKey/exhausted as a side effect, exactly
		// like the Python version's nonlocal mutations.
		_grouper := func(targetKey K) iter.Seq[V] {
			return func(yieldInner func(V) bool) {
				// Yield the current value first (it already matched).
				if !yieldInner(currValue) {
					return
				}

				// Continue pulling from the shared iterator.
				for {
					v, ok := next()
					if !ok {
						exhausted = true
						return
					}
					currValue = v
					currKey = keyFunc(currValue)

					if currKey != targetKey {
						// Key changed - stop this group. The outer loop
						// will pick up currValue/currKey for the next group.
						return
					}
					if !yieldInner(currValue) {
						return
					}
				}
			}
		}

		// Outer loop: emit (key, group) pairs until the input is exhausted.
		for !exhausted {
			targetKey := currKey
			currGroup := _grouper(targetKey)

			// Yield the key and its group to the caller.
			if !yield(targetKey, currGroup) {
				return
			}

			// If the caller didn't fully consume the group (i.e. broke out
			// of the inner loop early), we must drain it so that the shared
			// state advances past this group. This minnors the Python:
			//    if curr_key == target_key:
			//        for _ in curr_group: pass
			if currKey == targetKey {
				for range currGroup {
					// drain
				}
			}
		}
	}
}
