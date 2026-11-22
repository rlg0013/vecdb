package metrics

import "fmt"

func RecallAtK(
	groundTruth []int,
	retrieved []int,
	k int,
) (float64, error) {
	if k <= 0 {
		return 0, fmt.Errorf("k must be positive")
	}

	if len(groundTruth) < k {
		return 0, fmt.Errorf(
			"ground truth contains fewer than k elements",
		)
	}

	expected := make(map[int]struct{}, k)

	for _, id := range groundTruth[:k] {
		expected[id] = struct{}{}
	}

	found := make(map[int]struct{}, k)

	for _, id := range retrieved {
		if _, exists := expected[id]; exists {
			found[id] = struct{}{}
		}
	}

	recall := float64(len(found)) / float64(len(expected))

	return recall, nil
}
