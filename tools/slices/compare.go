package slices

import "math"

func CompareFloatSlices(a, b []float64, epsilon float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > epsilon {
			return false
		}
	}
	return true
}
