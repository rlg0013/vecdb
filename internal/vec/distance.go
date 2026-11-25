package vec

func checkDimensions(a, b []float32) {
	if len(a) != len(b) {
		panic("vector dimensions do not match")
	}
}

func L2Sq(a, b []float32) float32 {
	checkDimensions(a, b)

	var sum float32

	for i := 0; i < len(a); i++ {
		diff := a[i] - b[i]
		sum += diff * diff
	}

	return sum
}

func Dot(a, b []float32) float32 {
	checkDimensions(a, b)

	var sum float32

	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func CosineDist(a, b []float32) float32 {
	return 1.0 - Dot(a, b)
}
