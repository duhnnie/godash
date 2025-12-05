package godash

// FindFirst iterates over elements of the given slice and returns the first element
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
//   - fn: A function that takes an element of the slice and returns a boolean
//     indicating whether the element matches the criteria.
//
// Returns:
//   - The first element that matches the criteria and true, or the zero value of the
//     element type and false if no element matches.
func FindFirst[T any](slice []T, fn ElementIteratorFn[T, bool]) (T, bool, error) {
	for _, item := range slice {
		f, err := fn(item)
		if err != nil {
			return *new(T), false, err
		}

		if f {
			return item, true, nil
		}
	}

	return *new(T), false, nil
}

// FindFirstNE is the no-error variant of FindFirst. The provided iterator does not return an error
// and the function returns only the found element and a boolean indicator.
func FindFirstNE[T any](slice []T, fn ElementIteratorFnNE[T, bool]) (T, bool) {
	adapted := func(item T) (bool, error) {
		return fn(item), nil
	}

	r, found, _ := FindFirst(slice, adapted)
	return r, found
}
