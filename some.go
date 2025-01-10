package godash

func Some[T any](slice []T, iterator IteratorFn[T, bool]) bool {
	for index, item := range slice {
		if iterator(item, index, slice) {
			return true
		}
	}

	return false
}
