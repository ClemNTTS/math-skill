package lib

import "math"

func StandardDeviation(n float64) int {
	return int(math.Round(math.Sqrt(float64(n))))
}
