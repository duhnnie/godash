package godash

type IteratorFn[T, V any] func(currentValue T, currentIndex int, slice []T) V

func Every[T any](slice []T, iterator IteratorFn[T, bool]) bool {
	for index, item := range slice {
		if !iterator(item, index, slice) {
			return false
		}
	}

	return true
}
