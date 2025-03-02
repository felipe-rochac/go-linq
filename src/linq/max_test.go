package linq

import (
	"testing"
)

func TestLinq_Max(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[int]
		want int
	}{
		{
			name: "Empty Linq",
			l:    &Linq[int]{array: []int{}},
			want: 0,
		},
		{
			name: "Single element Linq",
			l:    &Linq[int]{array: []int{5}},
			want: 5,
		},
		{
			name: "Multiple elements Linq",
			l:    &Linq[int]{array: []int{1, 3, 2}},
			want: 3,
		},
		{
			name: "Negative elements Linq",
			l:    &Linq[int]{array: []int{-1, -3, -2}},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Max(); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_MaxF(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[float64]
		want float64
	}{
		{
			name: "Empty Linq",
			l:    &Linq[float64]{array: []float64{}},
			want: 0,
		},
		{
			name: "Single element Linq",
			l:    &Linq[float64]{array: []float64{5.5}},
			want: 5.5,
		},
		{
			name: "Multiple elements Linq",
			l:    &Linq[float64]{array: []float64{1.1, 3.3, 2.2}},
			want: 3.3,
		},
		{
			name: "Negative elements Linq",
			l:    &Linq[float64]{array: []float64{-1.1, -3.3, -2.2}},
			want: -1.1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.MaxF(); got != tt.want {
				t.Errorf("MaxF() = %v, want %v", got, tt.want)
			}
		})
	}
}
