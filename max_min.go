package godash

type numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Max[T numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T numeric](a, b T) T {
	if a < b {
		return a
	}
	return b
}
