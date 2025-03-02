package linq

// Average returns the average of the elements in the array.
func (l *Linq[T]) Average() int {
	return l.Sum() / l.Count()
}

// AverageF returns the average of the float64 elements in the array.
func (l *Linq[T]) AverageF() float64 {
	return l.SumF() / float64(l.Count())
}
