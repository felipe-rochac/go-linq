package linq

// Max returns the maximum element in the array.
func (l *Linq[T]) Max() int {
	if l.IsEmpty() || !isNumber(l.array[0]) {
		return 0
	}

	max := any(l.array[0]).(int)

	for _, item := range l.array {
		i := any(item).(int)
		if i > max {
			max = i
		}
	}
	return max
}

// MaxF returns the maximum float64 element in the array.
func (l *Linq[T]) MaxF() float64 {
	if l.IsEmpty() || !isNumber(l.array[0]) {
		return 0
	}

	max := any(l.array[0]).(float64)

	for _, item := range l.array {
		f := any(item).(float64)
		if f > max {
			max = f
		}
	}
	return max
}
