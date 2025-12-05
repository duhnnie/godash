package godash

// ReducerMapFn represents a function to execute for each element in the map during a reduce operation.
// The function receives the accumulator (result from the previous invocation or initial value),
// the current key, the current value, and the entire dictionary. It returns the updated accumulator
// and an error, if any.
//
// Parameters:
//
//	accumulator U: The accumulated value from the previous call, or the initial value on the first call.
//	currentKey K: The key of the current element in the map.
//	currentValue V: The value of the current element in the map.
//	dictionary map[K]V: The map being reduced.
//
// Returns:
//
//	U: The updated accumulator value.
//	error: An error, if any occurred during processing.
//
// Type Parameters:
//
//	K: The type of the map's keys (must be comparable).
//	V: The type of the map's values.
//	U: The type of the accumulator and return value.
type ReducerMapFn[K comparable, V, U any] func(accumulator U, currentKey K, currentValue V, dictionary map[K]V) (U, error)

// ReducerMapFnNE is a no-error variant of ReducerMapFn where the reducer does not return an error.
type ReducerMapFnNE[K comparable, V, U any] func(accumulator U, currentKey K, currentValue V, dictionary map[K]V) U

// ReduceMap applies a reducer function to each key-value pair in the provided map,
// accumulating a single result. The reducer function receives the current accumulator,
// the key and value of the current element, and the entire map. If the reducer returns
// an error, the reduction stops and the error is returned.
//
// Parameters:
//
//	m map[K]V: The map to reduce.
//	reducer ReducerMapFn[K, V, U]: The function to execute for each element.
//	initialValue U: The initial value for the accumulator.
//
// Returns:
//
//	U: The final accumulated value.
//	error: An error if any occurred during reduction.
//
// Type Parameters:
//
//	K: The type of the map's keys (must be comparable).
//	V: The type of the map's values.
//	U: The type of the accumulator and return value.
func ReduceMap[K comparable, V, U any](m map[K]V, reducer ReducerMapFn[K, V, U], initialValue U) (U, error) {
	var acc U = initialValue

	for key, value := range m {
		r, err := reducer(acc, key, value, m)
		if err != nil {
			return acc, err
		}

		acc = r
	}

	return acc, nil
}

// ReduceMapNE is the no-error variant of ReduceMap. It accepts a reducer that cannot return an error
// and returns only the accumulated value.
func ReduceMapNE[K comparable, V, U any](m map[K]V, reducer ReducerMapFnNE[K, V, U], initialValue U) U {
	adapted := func(acc U, key K, value V, dict map[K]V) (U, error) {
		return reducer(acc, key, value, dict), nil
	}

	r, _ := ReduceMap(m, adapted, initialValue)
	return r
}
