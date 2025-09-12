package slice

func Map[T, R any](source []T, resultingFunc func(item T) R) []R {
	result := make([]R, 0, len(source))

	for _, item := range source {
		result = append(result, resultingFunc(item))
	}

	return result
}
