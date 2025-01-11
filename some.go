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
func Some[T any](slice []T, iterator IteratorFn[T, bool]) bool {
	for index, item := range slice {
		if iterator(item, index, slice) {
			return true
		}
	}

	return false
}
