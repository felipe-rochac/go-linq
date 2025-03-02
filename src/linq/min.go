package linq

// Min returns the minimum element in the array.
func (l *Linq[T]) Min() int {
	if l.IsEmpty() || !isNumber(l.array[0]) {
		return 0
	}

	min := any(l.array[0]).(int)

	for _, item := range l.array {
		i := any(item).(int)
		if i < min {
			min = i
		}
	}
	return min
}

// MinF returns the minimum float64 element in the array.
func (l *Linq[T]) MinF() float64 {
	if l.IsEmpty() || !isNumber(l.array[0]) {
		return 0
	}

	min := any(l.array[0]).(float64)

	for _, item := range l.array {
		i := any(item).(float64)
		if i < min {
			min = i
		}
	}
	return min
}
