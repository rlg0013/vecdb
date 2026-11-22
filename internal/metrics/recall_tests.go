package metrics

import (
	"math"
	"testing"
)

func TestRecallAtK(t *testing.T) {
	tests := []struct {
		name        string
		groundTruth []int
		retrieved   []int
		k           int
		expected    float64
	}{
		{
			name:        "perfect recall",
			groundTruth: []int{1, 2, 3},
			retrieved:   []int{1, 2, 3},
			k:           3,
			expected:    1.0,
		},
		{
			name:        "partial recall",
			groundTruth: []int{1, 2, 3, 4, 5},
			retrieved:   []int{1, 2, 6, 7, 8},
			k:           5,
			expected:    0.4,
		},
		{
			name:        "zero recall",
			groundTruth: []int{1, 2, 3},
			retrieved:   []int{4, 5, 6},
			k:           3,
			expected:    0.0,
		},
		{
			name:        "extra retrieved results",
			groundTruth: []int{1, 2, 3},
			retrieved:   []int{4, 5, 1, 2, 3},
			k:           3,
			expected:    1.0,
		},
		{
			name:        "duplicate retrieved IDs",
			groundTruth: []int{1, 2, 3},
			retrieved:   []int{1, 1, 1, 2},
			k:           3,
			expected:    2.0 / 3.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RecallAtK(
				tt.groundTruth,
				tt.retrieved,
				tt.k,
			)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf(
					"expected recall %f, got %f",
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestRecallAtKInvalidK(t *testing.T) {
	_, err := RecallAtK(
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		0,
	)

	if err == nil {
		t.Fatal("expected error for k = 0")
	}
}

func TestRecallAtKInsufficientGroundTruth(t *testing.T) {
	_, err := RecallAtK(
		[]int{1, 2},
		[]int{1, 2, 3},
		3,
	)

	if err == nil {
		t.Fatal("expected error for insufficient ground truth")
	}
}
