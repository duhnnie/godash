package godash

// Filter iterates over elements of the provided slice and returns a new slice
// containing all elements for which the provided function returns true.
//
// The function takes two parameters:
//   - slice: A slice of any type T.
//   - fn: A function that takes an element of type T and returns a boolean indicating
//     whether the element should be included in the result. The function may also return
//     an error if the evaluation fails.
//
// Example usage:
//
//	result, err := Filter([]int{1, 2, 3, 4}, func(item int) (bool, error) {
//	    return item%2 == 0, nil
//	})
//	// result will be []int{2, 4}
//
// Type Parameters:
// - T: The type of elements in the slice.
//
// Returns:
// A slice of type T containing all elements for which the provided function returns true.
// An error is returned if the predicate function returns an error during evaluation.
func Filter[T any](slice []T, fn ElementIteratorFn[T, bool]) ([]T, error) {
	var r []T

	for _, item := range slice {
		f, err := fn(item)
		if err != nil {
			return nil, err
		}

		if f {
			r = append(r, item)
		}
	}

	return r, nil
}

// FilterNE is the no-error variant of Filter. It accepts an element-only predicate
// that does not return an error and returns the filtered slice.
func FilterNE[T any](slice []T, fn ElementIteratorFnNE[T, bool]) []T {
	adapted := func(item T) (bool, error) {
		return fn(item), nil
	}

	r, _ := Filter(slice, adapted)
	return r
}
