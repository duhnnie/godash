package godash

// FindAll iterates over elements of the provided slice and returns a new slice
// containing all elements for which the provided function returns true.
//
// The function takes two parameters:
//   - slice: A slice of any type T.
//   - fn: A function that takes an element of type T, its index, and the original slice,
//     and returns a boolean indicating whether the element should be included in the result.
//     The function may also return an error if the evaluation fails.
//
// Example usage:
//
//	result, err := FindAll([]int{1, 2, 3, 4}, func(item int, index int, slice []int) (bool, error) {
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
func FindAll[T any](slice []T, fn IteratorFn[T, bool]) ([]T, error) {
	var r []T

	for index, item := range slice {
		f, err := fn(item, index, slice)
		if err != nil {
			return nil, err
		}

		if f {
			r = append(r, item)
		}
	}

	return r, nil
}

// FindAllNE is the no-error variant of FindAll. The provided iterator does not return an error
// and the function returns only the resulting slice.
func FindAllNE[T any](slice []T, fn IteratorFnNE[T, bool]) []T {
	adapted := func(item T, idx int, s []T) (bool, error) {
		return fn(item, idx, s), nil
	}

	r, _ := FindAll(slice, adapted)
	return r
}
