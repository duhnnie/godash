package godash

// FindAll iterates over elements of the provided slice and returns a new slice
// containing all elements for which the provided function returns true.
//
// The function takes two parameters:
//   - slice: A slice of any type T.
//   - fn: A function that takes an element of type T, its index, and the original slice,
//     and returns a boolean indicating whether the element should be included in the result.
//
// Example usage:
//
//	result := FindAll([]int{1, 2, 3, 4}, func(item int, index int, slice []int) bool {
//	    return item%2 == 0
//	})
//	// result will be []int{2, 4}
//
// Type Parameters:
// - T: The type of elements in the slice.
//
// Returns:
// A slice of type T containing all elements for which the provided function returns true.
func FindAll[T any](slice []T, fn IteratorFn[T, bool]) []T {
	var r []T

	for index, item := range slice {
		if fn(item, index, slice) {
			r = append(r, item)
		}
	}

	return r
}
