package linq

import (
	"testing"
)

func TestLinq_Sum(t *testing.T) {
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
			name: "Non-empty Linq",
			l:    &Linq[int]{array: []int{1, 2, 3}},
			want: 6,
		},
		{
			name: "Negative numbers",
			l:    &Linq[int]{array: []int{-1, -2, -3}},
			want: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_SumF(t *testing.T) {
	tests := []struct {
		name string
		l    *Linq[float64]
		want float64
	}{
		{
			name: "Empty Linq",
			l:    &Linq[float64]{array: []float64{}},
			want: 0.0,
		},
		{
			name: "Non-empty Linq",
			l:    &Linq[float64]{array: []float64{1.1, 2.2, 3.3}},
			want: 6.6,
		},
		{
			name: "Negative numbers",
			l:    &Linq[float64]{array: []float64{-1.1, -2.2, -3.3}},
			want: -6.6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.SumF(); got != tt.want {
				t.Errorf("SumF() = %v, want %v", got, tt.want)
			}
		})
	}
}
