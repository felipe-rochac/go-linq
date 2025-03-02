package linq

import (
	"testing"
)

func TestLinq_Min(t *testing.T) {
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
			l:    &Linq[int]{array: []int{3, 1, 4, 1, 5, 9}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Min(); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinq_MinF(t *testing.T) {
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
			l:    &Linq[float64]{array: []float64{3.1, 1.4, 4.1, 1.5, 5.9, 9.2}},
			want: 1.4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.MinF(); got != tt.want {
				t.Errorf("MinF() = %v, want %v", got, tt.want)
			}
		})
	}
}
