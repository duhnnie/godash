// Package godash provides a collection of utility functions for working with slices and maps,
// inspired by functional programming concepts similar to Lodash for JavaScript.
package godash

// ElementIteratorFn is a function type that transforms a single element of type T into a value of type U,
// returning an error if the transformation fails.
//
// This is a simpler alternative to MapFn, accepting only the element value without index
// or collection reference information.
//
// Type parameters:
//   - T: the input element type
//   - U: the output element type
//
// Parameters:
//   - element: the value to transform
//
// Returns:
//   - U: the transformed value
//   - error: an error if the transformation fails, nil otherwise
type ElementIteratorFn[T any, U any] func(element T) (U, error)

// ElementIteratorFnNE is a no-error variant of ElementIteratorFn that does not return an error.
type ElementIteratorFnNE[T any, U any] func(element T) U

// MapTo applies a transformation function to each element in a collection and returns a new slice
// containing the transformed values. If any transformation fails, the operation is aborted and
// the error is returned.
//
// This function is useful for converting a slice of one type to another type, with error handling
// for each transformation.
//
// Type parameters:
//   - T: the element type of the input collection
//   - U: the element type of the output collection
//
// Parameters:
//   - collection: the input slice to transform
//   - mapFunction: the transformation function to apply to each element
//
// Returns:
//   - []U: a new slice containing the transformed elements
//   - error: an error if any transformation fails, nil otherwise
func MapTo[T any, U any](collection []T, mapFunction ElementIteratorFn[T, U]) ([]U, error) {
	mapped := make([]U, len(collection))

	for index, element := range collection {
		m, err := mapFunction(element)
		if err != nil {
			return mapped, err
		}

		mapped[index] = m
	}

	return mapped, nil
}

// MapToNE is a convenience wrapper for MapTo that accepts a no-error mapper and
// returns the transformed slice without an error return value (errors from the
// underlying MapTo are ignored since the provided mapper cannot return an error).
func MapToNE[T any, U any](collection []T, mapFunction ElementIteratorFnNE[T, U]) []U {
	adapted := func(element T) (U, error) {
		return mapFunction(element), nil
	}

	r, _ := MapTo(collection, adapted)
	return r
}
