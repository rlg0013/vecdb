package vec

import (
	"math"
	"testing"
)

func almostEqual(a, b float32) bool {
	const epsilon = 1e-5
	return math.Abs(float64(a-b)) < epsilon
}

func TestL2Sq(t *testing.T) {
	tests := []struct {
		name string
		a    []float32
		b    []float32
		want float32
	}{
		{
			name: "basic vectors",
			a:    []float32{2, 3},
			b:    []float32{4, 1},
			want: 8,
		},
		{
			name: "identical vectors",
			a:    []float32{1, 2, 3},
			b:    []float32{1, 2, 3},
			want: 0,
		},
		{
			name: "opposite vectors",
			a:    []float32{1, 2},
			b:    []float32{-1, -2},
			want: 20,
		},
		{
			name: "empty vectors",
			a:    []float32{},
			b:    []float32{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := L2Sq(tt.a, tt.b)

			if !almostEqual(got, tt.want) {
				t.Errorf("L2Sq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		name string
		a    []float32
		b    []float32
		want float32
	}{
		{
			name: "basic vectors",
			a:    []float32{2, 3},
			b:    []float32{4, 1},
			want: 11,
		},
		{
			name: "orthogonal vectors",
			a:    []float32{1, 0},
			b:    []float32{0, 1},
			want: 0,
		},
		{
			name: "opposite vectors",
			a:    []float32{1, 2},
			b:    []float32{-1, -2},
			want: -5,
		},
		{
			name: "empty vectors",
			a:    []float32{},
			b:    []float32{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dot(tt.a, tt.b)

			if !almostEqual(got, tt.want) {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCosineDist(t *testing.T) {
	tests := []struct {
		name string
		a    []float32
		b    []float32
		want float32
	}{
		{
			name: "same direction",
			a:    []float32{1, 0},
			b:    []float32{1, 0},
			want: 0,
		},
		{
			name: "perpendicular vectors",
			a:    []float32{1, 0},
			b:    []float32{0, 1},
			want: 1,
		},
		{
			name: "opposite direction",
			a:    []float32{1, 0},
			b:    []float32{-1, 0},
			want: 2,
		},
		{
			name: "diagonal normalized vectors",
			a:    []float32{0.6, 0.8},
			b:    []float32{0.8, 0.6},
			want: 0.04,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CosineDist(tt.a, tt.b)

			if !almostEqual(got, tt.want) {
				t.Errorf("CosineDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDimensionMismatch(t *testing.T) {
	a := []float32{1, 2}
	b := []float32{1, 2, 3}

	tests := []struct {
		name string
		fn   func()
	}{
		{
			name: "L2Sq",
			fn: func() {
				L2Sq(a, b)
			},
		},
		{
			name: "Dot",
			fn: func() {
				Dot(a, b)
			},
		},
		{
			name: "CosineDist",
			fn: func() {
				CosineDist(a, b)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("expected panic for dimension mismatch")
				}
			}()

			tt.fn()
		})
	}
}
