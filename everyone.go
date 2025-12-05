package godash

// Everyone iterates over elements of a slice, returning true if all elements pass the
// test implemented by the provided iterator function. It stops iterating as soon as
// the iterator function returns false or an error for any element.
//
// Parameters:
//   - slice: The slice to iterate over.
//   - iterator: A function that is invoked for each element in the slice. It takes
//     the current element and returns a boolean indicating whether the element passes
//     the test, along with a potential error.
//
// Returns:
//   - bool: true if all elements pass the test (iterator returns true for each element),
//     false otherwise.
//   - error: An error if the iterator function encounters an error during iteration,
//     nil otherwise.
//
// Example:
//
//	slice := []int{2, 4, 6, 8}
//	allEven, err := Everyone(slice, func(num int) (bool, error) {
//		return num%2 == 0, nil
//	})
//	// allEven is true
func Everyone[T any](slice []T, iterator ElementIteratorFn[T, bool]) (bool, error) {
	for _, item := range slice {
		r, err := iterator(item)
		if err != nil {
			return false, err
		}

		if !r {
			return false, nil
		}
	}

	return true, nil
}
