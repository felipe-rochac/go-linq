package linq

import (
	"reflect"
	"testing"
)

func TestLinq_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[int]
		want []int
	}{
		{
			name: "Empty Linq",
			l:    &Linq[int]{array: []int{}},
			want: []int{},
		},
		{
			name: "Non-empty Linq",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	tests := []struct {
		name           string
		l              *Linq[int]
		keyPredicate   func(int) any
		valuePredicate func(int) any
		want           map[any]any
	}{
		{
			name: "Empty Linq",
			l:    &Linq[int]{array: []int{}},
			keyPredicate: func(i int) any {
				return i
			},
			valuePredicate: func(i int) any {
				return ""
			},
			want: map[any]any{},
		},
		{
			name: "Non-empty Linq",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			keyPredicate: func(i int) any {
				return i
			},
			valuePredicate: func(i int) any {
				return string(rune('a' + i - 1))
			},
			want: map[any]any{1: "a", 2: "b", 3: "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.ToMap(tt.keyPredicate, tt.valuePredicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
