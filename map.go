package godash

// MapFn is a function that receives:
//
// element T: a value of any type.
//
// index int: the index in the target collection of the element T.
//
// collection []T: the collection that contains the element provided.
//
// It returns a value of any type.
type MapFn[T, U any] func(element T, index int, collection []T) U

// Map creates a new slice populated with the results of calling a provided
// function on every element in the calling slice.
//
// collection is a slice of any type.
//
// mapFunction is the function to call on every element in the calling slice.
//
// It returns a slice of any type.
func Map[T, U any](collection []T, mapFunction MapFn[T, U]) []U {
	mapped := make([]U, len(collection))

	for index, element := range collection {
		mapped[index] = mapFunction(element, index, collection)
	}

	return mapped
}
