package godash

// Any iterates over elements of the provided slice and invokes the given iterator function
// for each element. If the iterator function returns true for any element, Any returns true.
// If the iterator function never returns true, Any returns false.
//
// T represents the type of elements in the slice.
//
// Parameters:
//   - slice: The slice to iterate over.
//   - iterator: A function that takes an element of the slice and returns a boolean
//     indicating whether the condition is met.
//
// Returns:
//   - bool: true if the iterator function returns true for any element, otherwise false.
func Any[T any](slice []T, iterator ElementIteratorFn[T, bool]) (bool, error) {
	for _, item := range slice {
		f, err := iterator(item)
		if err != nil {
			return false, err
		}

		if f {
			return true, nil
		}
	}

	return false, nil
}

// AnyNE is the no-error variant of Any. The provided iterator does not return an error
// and AnyNE returns only the boolean result.
func AnyNE[T any](slice []T, iterator ElementIteratorFnNE[T, bool]) bool {
	adapted := func(element T) (bool, error) {
		return iterator(element), nil
	}

	r, _ := Any(slice, adapted)
	return r
}
