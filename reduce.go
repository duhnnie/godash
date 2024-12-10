package godash

// ReduceFn represents a function to execute for each element in the slice. Its
// return value becomes the value of the accumulator parameter on the next
// invocation of callbackFn. For the last invocation, the return value becomes
// the return value of reduce(). The function is called with the following
// arguments:
//
// accumulator U: The value resulting from the previous call to callbackFn. On
// the first call, its value is initialValue if the latter is specified;
// otherwise its value is slice[0].
//
// currentValue T: The value of the current element. On the first call, its
// value is slice[0] if initialValue is specified; otherwise its value is
// slice[1].
//
// currentIndex int: The index position of currentValue in the slice. On the
// first call, its value is 0 and it's incremented for each next call.
//
// slice []T: The slice reduce() was called upon.
type ReducerFn[T any, U any] func(accumulator U, currentValue T, currentIndex int, slice []T) U

// Reduce executes a user-supplied "reducer" callback function on each element
// of the slice, in order, passing in the return value from the calculation on
// the preceding element. The final result of running the reducer across all
// elements of the slice is a single value. It receives the following
// arguments:
//
// slice []T: the slice to reduce.
//
// reducer ReducerFn[T, U]: The reducer function.
//
// initialValue U: A value to which accumulator is initialized the first time
// the reducer function is called.
func Reduce[T any, U any](slice []T, reducer ReducerFn[T, U], initialValue U) U {
	var acc U = initialValue

	for index, item := range slice {
		acc = reducer(acc, item, index, slice)
	}

	return acc
}
