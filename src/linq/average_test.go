package linq

import (
	"testing"
)

func TestLinq_Average(t *testing.T) {
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
			l:    &Linq[int]{array: []int{1, 2, 3, 4, 5}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Average(); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_AverageF(t *testing.T) {
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
			l:    &Linq[float64]{array: []float64{5.0}},
			want: 5.0,
		},
		{
			name: "Multiple elements Linq",
			l:    &Linq[float64]{array: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
			want: 3.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.AverageF(); got != tt.want {
				t.Errorf("AverageF() = %v, want %v", got, tt.want)
			}
		})
	}
}
