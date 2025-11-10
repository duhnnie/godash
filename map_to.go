package godash

// MapToFn is a simpler version of MapFn that only receives:
//
// element T: a value of any type.
//
// It returns a value of any type.
type MapToFn[T any, U any] func(element T) U

// MapTo is similar to Map but uses a simpler mapping function that only receives
// the element value, without the index or collection reference. It creates a new
// slice populated with the results of calling the provided function on every
// element in the input slice.
//
// collection is a slice of any type.
//
// mapFunction is the function to call on every element, receiving only the element value.
//
// It returns a slice of any type.
func MapTo[T any, U any](collection []T, mapFunction MapToFn[T, U]) []U {
	mapped := make([]U, len(collection))

	for index, element := range collection {
		mapped[index] = mapFunction(element)
	}

	return mapped
}
