package linq

type Linq[T any] struct {
	array []T
}

// From creates a new Linq instance from the given array.
func From[T any](array []T) *Linq[T] {
	return &Linq[T]{array: array}
}

// isNumber checks if the given type is a Number.
func isNumber(t any) bool {
	switch t.(type) {
	case int, int32, int64, float32, float64, uint, uint32, uint64:
		return true
	default:
		return false
	}
}

func (l *Linq[T]) IsEmpty() bool {
	return len(l.array) == 0
}
