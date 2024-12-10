package godash

type MapFn[T, U any] func(element T, index int, collection []T) U

func Map[T, U any](collection []T, mapFunction MapFn[T, U]) []U {
	mapped := make([]U, len(collection))

	for index, element := range collection {
		mapped[index] = mapFunction(element, index, collection)
	}

	return mapped
}
