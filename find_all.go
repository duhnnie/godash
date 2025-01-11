package godash

func FindAll[T any](slice []T, fn IteratorFn[T, bool]) []T {
	var r []T

	for index, item := range slice {
		if fn(item, index, slice) {
			r = append(r, item)
		}
	}

	return r
}
