package linq

import (
	"reflect"
	"testing"
)

func TestFrom(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  *Linq[int]
	}{
		{
			name:  "Empty array",
			input: []int{},
			want:  &Linq[int]{array: []int{}},
		},
		{
			name:  "Non-empty array",
			input: []int{1, 2, 3},
			want:  &Linq[int]{array: []int{1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := From(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("From() = %v, want %v", got, tt.want)
			}
		})
	}
}
