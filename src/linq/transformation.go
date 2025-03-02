package linq

// ToArray returns the array of the Linq instance.
func (l *Linq[T]) ToSlice() []T {
	return l.array
}

// ToSlice returns the slice of the Linq instance.
func (l *Linq[T]) ToMap(keyPredicate func(T) any, valuePredicate func(T) any) map[any]any {
	result := make(map[any]any)
	for _, item := range l.array {
		key := keyPredicate(item)
		result[key] = valuePredicate(item)
	}
	return result
}
