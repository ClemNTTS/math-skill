package lib

func Average(t []int) float64 {
	var Somme int
	for _, n := range t {
		Somme += n
	}
	return float64(Somme) / float64(len(t))
}
