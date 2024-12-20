package godash

// ReducerMapFn represents a function to execute for each element in the map. Its
// return value becomes the value of the accumulator parameter on the next
// invocation of callbackFn. For the last invocation, the return value becomes
// the return value of reduce(). The function is called with the following
// arguments:
//
// accumulator U: The value resulting from the previous call to callbackFn. On
// the first call, its value is initialValue if the latter is specified;
// otherwise its value is dictionary[firstKey].
//
// currentKey K: The key of the current element.
//
// currentValue V: The value of the current element.
//
// dictionary map[K]V: The map reduce() was called upon.
type ReducerMapFn[K comparable, V, U any] func(accumulator U, currentKey K, currentValue V, dictionary map[K]V) U

// ReduceMap executes a user-supplied "reducer" callback function on each element
// of the map, in order, passing in the return value from the calculation on
// the preceding element. The final result of running the reducer across all
// elements of the map is a single value. It receives the following
// arguments:
//
// m map[K]V: the map to reduce.
//
// reducer ReducerMapFn[K, V, U]: The reducer function.
//
// initialValue U: A value to which accumulator is initialized the first time
// the reducer function is called.
func ReduceMap[K comparable, V, U any](m map[K]V, reducer ReducerMapFn[K, V, U], initialValue U) U {
	var acc U = initialValue

	for key, value := range m {
		acc = reducer(acc, key, value, m)
	}

	return acc
}
