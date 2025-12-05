package godash

// ReducerFn is a callback function type used by the Reduce function to process
// each element in a slice and accumulate a result.
//
// The function receives four parameters:
//   - accumulator: The accumulated value from the previous iteration, or initialValue
//     on the first call
//   - currentValue: The current element being processed from the slice
//   - currentIndex: The zero-based index of the current element in the slice
//   - slice: The original slice being reduced
//
// The function returns the new accumulator value and an error. If an error is
// returned, the Reduce operation is immediately terminated and the error is
// propagated to the caller.
//
// Type parameters:
//   - T: The type of elements in the slice being reduced
//   - U: The type of the accumulator and the final result value
type ReducerFn[T, U any] func(accumulator U, currentValue T, currentIndex int, slice []T) (U, error)

// ReducerFnNE is a no-error variant of ReducerFn where the reducer does not return an error.
type ReducerFnNE[T, U any] func(accumulator U, currentValue T, currentIndex int, slice []T) U

// Reduce applies a reducer function against an accumulator and each element in
// a slice (from left to right), resulting in a single value. The reducer is
// invoked for each element in the slice, receiving the accumulated result from
// the previous invocation.
//
// Parameters:
//   - slice: The slice to reduce. If empty, initialValue is returned.
//   - reducer: The function to execute on each element to update the accumulator.
//   - initialValue: The initial value for the accumulator on the first call.
//
// Returns:
//   - The final accumulated value after processing all elements.
//   - An error if the reducer function returns an error for any element.
//
// Type parameters:
//   - T: The type of elements in the input slice
//   - U: The type of the accumulator and result value (may differ from T)
func Reduce[T any, U any](slice []T, reducer ReducerFn[T, U], initialValue U) (U, error) {
	var acc U = initialValue

	for index, item := range slice {
		r, err := reducer(acc, item, index, slice)
		if err != nil {
			return acc, err
		}

		acc = r
	}

	return acc, nil
}

// ReduceNE is the no-error variant of Reduce. It accepts a reducer that cannot return an error
// and returns the accumulated value directly.
func ReduceNE[T any, U any](slice []T, reducer ReducerFnNE[T, U], initialValue U) U {
	adapted := func(acc U, cur T, idx int, s []T) (U, error) {
		return reducer(acc, cur, idx, s), nil
	}

	r, _ := Reduce(slice, adapted, initialValue)
	return r
}
