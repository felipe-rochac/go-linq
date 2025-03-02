package linq

import (
	"reflect"
	"testing"
)

func TestLinq_Where(t *testing.T) {
	tests := []struct {
		name      string
		l         *Linq[int]
		selector  func(int) bool
		wantArray []int
	}{
		{
			name: "Filter even numbers",
			l:    &Linq[int]{array: []int{1, 2, 3, 4, 5}},
			selector: func(i int) bool {
				return i%2 == 0
			},
			wantArray: []int{2, 4},
		},
		{
			name: "Filter greater than 3",
			l:    &Linq[int]{array: []int{1, 2, 3, 4, 5}},
			selector: func(i int) bool {
				return i > 3
			},
			wantArray: []int{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Where(tt.selector); !reflect.DeepEqual(got.array, tt.wantArray) {
				t.Errorf("Where() = %v, want %v", got.array, tt.wantArray)
			}
		})
	}
}

func TestLinq_Select(t *testing.T) {
	tests := []struct {
		name      string
		l         *Linq[int]
		selector  func(int) any
		wantArray []any
	}{
		{
			name: "Select squares",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			selector: func(i int) any {
				return i * i
			},
			wantArray: []any{1, 4, 9},
		},
		{
			name: "Select strings",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			selector: func(i int) any {
				return string(rune('a' + i - 1))
			},
			wantArray: []any{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Select(tt.selector); !reflect.DeepEqual(got.array, tt.wantArray) {
				t.Errorf("Select() = %v, want %v", got.array, tt.wantArray)
			}
		})
	}
}

func TestLinq_Distinct(t *testing.T) {
	tests := []struct {
		name      string
		l         *Linq[int]
		wantArray []int
	}{
		{
			name:      "Distinct elements",
			l:         &Linq[int]{array: []int{1, 2, 2, 3, 3, 3}},
			wantArray: []int{1, 2, 3},
		},
		{
			name:      "All distinct",
			l:         &Linq[int]{array: []int{1, 2, 3}},
			wantArray: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Distinct(); !reflect.DeepEqual(got.array, tt.wantArray) {
				t.Errorf("Distinct() = %v, want %v", got.array, tt.wantArray)
			}
		})
	}
}

func TestLinq_First(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[int]
		want int
	}{
		{
			name: "First element",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			want: 1,
		},
		{
			name: "Single element",
			l:    &Linq[int]{array: []int{42}},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.First(); got != tt.want {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_Last(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[int]
		want int
	}{
		{
			name: "Last element",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			want: 3,
		},
		{
			name: "Single element",
			l:    &Linq[int]{array: []int{42}},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Last(); got != tt.want {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_Count(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[int]
		want int
	}{
		{
			name: "Count elements",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			want: 3,
		},
		{
			name: "Empty array",
			l:    &Linq[int]{array: []int{}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_Any(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[int]
		want bool
	}{
		{
			name: "Non-empty array",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			want: true,
		},
		{
			name: "Empty array",
			l:    &Linq[int]{array: []int{}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Any(); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_Take(t *testing.T) {
	tests := []struct {
		name  string
		l     *Linq[int]
		from  int
		count int
		want  []int
	}{
		{
			name:  "Take first 2 elements",
			l:     &Linq[int]{array: []int{1, 2, 3, 4, 5}},
			from:  0,
			count: 2,
			want:  []int{1, 2},
		},
		{
			name:  "Take middle 2 elements",
			l:     &Linq[int]{array: []int{1, 2, 3, 4, 5}},
			from:  2,
			count: 2,
			want:  []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Take(tt.from, tt.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_All(t *testing.T) {
	tests := []struct {
		name      string
		l         *Linq[int]
		predicate func(int) bool
		want      bool
	}{
		{
			name: "All even",
			l:    &Linq[int]{array: []int{2, 4, 6}},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: true,
		},
		{
			name: "Not all even",
			l:    &Linq[int]{array: []int{2, 3, 6}},
			predicate: func(i int) bool {
				return i%2 == 0
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.All(tt.predicate); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
