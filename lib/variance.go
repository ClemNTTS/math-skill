package lib

func Variance(moy int, tab []int) int {
	var vari float64
	for _, value := range tab {
		vari = vari + RecursivePower((float64(value)-float64(moy)), 2)
	}
	return int(vari / float64(len(tab)))
}

func RecursivePower(nb float64, power float64) float64 {
	cpt := power
	if power < 0 {
		return 0
	}
	if cpt != 0 {
		cpt--
		return RecursivePower(nb, cpt) * nb
	}
	return 1
}
