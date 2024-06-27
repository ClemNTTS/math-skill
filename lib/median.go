package lib

func Median(t []int) float64 {
	if len(t)%2 == 1 {
		return float64(t[len(t)/2])
	} else {
		return ((float64(t[len(t)/2]) + float64(t[len(t)/2-1])) / float64(2))
	}
}
