package godash

// IteratorFn is a generic function type that represents an iterator function.
// It takes the following parameters:
// - currentValue: The current element being processed in the slice.
// - currentIndex: The index of the current element being processed in the slice.
// - slice: The slice that is being iterated over.
// It returns a value of type V and a potential error.
type IteratorFn[T, V any] func(currentValue T, currentIndex int, slice []T) (V, error)

// IteratorFnNE is the no-error variant of IteratorFn where the callback does not return an error.
type IteratorFnNE[T, V any] func(currentValue T, currentIndex int, slice []T) V

// Every iterates over elements of a slice, returning true if all elements pass the
// test implemented by the provided iterator function. It stops iterating as soon as
// the iterator function returns false for any element.
//
// This function implements the behavior of Array.every() from JavaScript, allowing
// you to test whether all elements in a slice satisfy a provided testing function.
// If any iterator invocation returns an error, iteration stops immediately and the
// error is returned.
//
// Parameters:
//   - slice: The slice to iterate over.
//   - iterator: A function that is invoked for each element in the slice. It takes
//     the current element, its index, and the entire slice as arguments, and returns
//     a boolean indicating whether the element passes the test, along with a potential error.
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
//	allEven, err := Every(slice, func(num int, idx int, s []int) (bool, error) {
//		return num%2 == 0, nil
//	})
//	// allEven is true
func Every[T any](slice []T, iterator IteratorFn[T, bool]) (bool, error) {
	for index, item := range slice {
		r, err := iterator(item, index, slice)
		if err != nil {
			return false, err
		}

		if !r {
			return false, nil
		}
	}

	return true, nil
}

// EveryoneNE is the no-error variant of Everyone which accepts a callback that does not return an error.
func EveryoneNE[T any](slice []T, iterator ElementIteratorFnNE[T, bool]) bool {
	adapted := func(v T) (bool, error) {
		return iterator(v), nil
	}

	r, _ := Everyone(slice, adapted)
	return r
}
