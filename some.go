package godash

// Some iterates over elements of the provided slice and invokes the given iterator function
// for each element. If the iterator function returns true for any element, Some returns true.
// If the iterator function never returns true, Some returns false.
//
// T represents the type of elements in the slice.
//
// Parameters:
//   - slice: The slice to iterate over.
//   - iterator: A function that takes an element of the slice, its index, and the entire slice,
//     and returns a boolean indicating whether the condition is met.
//
// Returns:
//   - bool: true if the iterator function returns true for any element, otherwise false.
func Some[T any](slice []T, iterator IteratorFn[T, bool]) (bool, error) {
	for index, item := range slice {
		f, err := iterator(item, index, slice)
		if err != nil {
			return false, err
		}

		if f {
			return true, nil
		}
	}

	return false, nil
}

// SomeNE is the no-error variant of Some. It accepts a callback that does not return an error
// and returns only the boolean result.
func SomeNE[T any](slice []T, iterator IteratorFnNE[T, bool]) bool {
	adapted := func(v T, idx int, s []T) (bool, error) {
		return iterator(v, idx, s), nil
	}

	r, _ := Some(slice, adapted)
	return r
}
