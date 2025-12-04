package godash

// MapFn is a function type that transforms elements of a collection.
//
// Parameters:
//   - element T: The current element being processed.
//   - index int: The zero-based index of the current element in the collection.
//   - collection []T: The entire collection being iterated over.
//
// Returns:
//   - U: The transformed value.
//   - error: An error if the transformation fails.
type MapFn[T, U any] func(element T, index int, collection []T) (U, error)

// Map transforms each element of a collection using the provided function
// and returns a new slice containing the transformed elements.
//
// Parameters:
//   - collection: The input slice of elements to transform.
//   - mapFunction: The transformation function to apply to each element.
//
// Returns:
//   - A new slice containing the transformed elements.
//   - An error if any transformation fails. The partially populated result
//     is returned along with the error.
func Map[T, U any](collection []T, mapFunction MapFn[T, U]) ([]U, error) {
	mapped := make([]U, len(collection))

	for index, element := range collection {
		m, err := mapFunction(element, index, collection)
		if err != nil {
			return mapped, err
		}

		mapped[index] = m
	}

	return mapped, nil
}
