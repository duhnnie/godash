package godash

func Find[T any](slice []T, fn IteratorFn[T, bool]) (T, bool) {
	for index, item := range slice {
		if fn(item, index, slice) {
			return item, true
		}
	}

	var x T

	return x, false
}
