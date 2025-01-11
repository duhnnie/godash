package godash

// IteratorFn is a generic function type that represents an iterator function.
// It takes the following parameters:
// - currentValue: The current element being processed in the slice.
// - currentIndex: The index of the current element being processed in the slice.
// - slice: The slice that is being iterated over.
// It returns a value of type V.
type IteratorFn[T, V any] func(currentValue T, currentIndex int, slice []T) V

// Every iterates over elements of a slice, returning true if all elements pass the
// test implemented by the provided iterator function. It stops iterating as soon as
// the iterator function returns false for any element.
//
// Parameters:
//   - slice: The slice to iterate over.
//   - iterator: A function that is invoked for each element in the slice. It takes
//     the current element, its index, and the entire slice as arguments, and returns
//     a boolean indicating whether the element passes the test.
//
// Returns:
// - bool: true if all elements pass the test, otherwise false.
func Every[T any](slice []T, iterator IteratorFn[T, bool]) bool {
	for index, item := range slice {
		if !iterator(item, index, slice) {
			return false
		}
	}

	return true
}
