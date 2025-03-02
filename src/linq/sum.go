package linq

// Sum returns the sum of the elements in the array.
func (l *Linq[T]) Sum() int {
	if l.IsEmpty() || !isNumber(l.array[0]) {
		return 0
	}

	var sum int
	for _, item := range l.array {
		i := any(item).(int)
		sum += i
	}
	return sum
}

// SumF returns the sum of the float64 elements in the array.
func (l *Linq[T]) SumF() float64 {
	var sum float64
	for _, item := range l.array {
		i := any(item).(float64)
		sum += i
	}
	return sum
}
