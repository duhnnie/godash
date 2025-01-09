package godash

func Clamp[T numeric](v T, min T, max T) T {
	return Max(Min(v, max), min)
}
