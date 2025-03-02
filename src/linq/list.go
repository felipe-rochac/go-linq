package linq

// Where filters the elements of the array based on the given predicate.
func (l *Linq[T]) Where(selector func(T) bool) *Linq[T] {
	var result []T
	for _, item := range l.array {
		if selector(item) {
			result = append(result, item)
		}
	}
	l.array = result
	return l
}

// Select projects each element of the array into a new form.
func (l *Linq[T]) Select(selector func(T) any) *Linq[any] {
	result := make([]any, 0, len(l.array))
	for _, item := range l.array {
		result = append(result, selector(item))
	}
	return &Linq[any]{array: result}
}

// Distinct returns the distinct elements in the array.
func (l *Linq[T]) Distinct() *Linq[T] {
	seen := make(map[any]struct{})
	var result []T
	for _, item := range l.array {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	l.array = result
	return l
}

// First returns the first element of the array.
func (l *Linq[T]) First() T {
	return l.array[0]
}

// Last returns the last element of the array.
func (l *Linq[T]) Last() T {
	return l.array[len(l.array)-1]
}

// Count returns the number of elements in the array.
func (l *Linq[T]) Count() int {
	return len(l.array)
}

// Any returns true if the array has any elements.
func (l *Linq[T]) Any() bool {
	return len(l.array) > 0
}

// Take returns the elements from the array starting from the given index and count.
func (l *Linq[T]) Take(from, count int) []T {
	return l.array[from : from+count]
}

// All returns true if all elements in the array satisfy the given predicate.
func (l *Linq[T]) All(predicate func(T) bool) bool {
	for _, item := range l.array {
		if !predicate(item) {
			return false
		}
	}
	return true
}
