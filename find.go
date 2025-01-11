package godash

// Find iterates over elements of the given slice and returns the first element
// for which the provided function returns true. The function takes an element,
// its index, and the entire slice as arguments.
//
// The function returns the found element and true if such an element is found.
// If no element matches the criteria, it returns the zero value of the element
// type and false.
//
// T is a generic type parameter that can be any type.
//
// Parameters:
//   - slice: The slice to iterate over.
//   - fn: A function that takes an element of the slice, its index, and the entire
//     slice, and returns a boolean indicating whether the element matches the criteria.
//
// Returns:
//   - The first element that matches the criteria and true, or the zero value of the
//     element type and false if no element matches.
func Find[T any](slice []T, fn IteratorFn[T, bool]) (T, bool) {
	for index, item := range slice {
		if fn(item, index, slice) {
			return item, true
		}
	}

	var x T

	return x, false
}
