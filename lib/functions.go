package lib

import (
	"os"
	"strconv"
)

// Scrap all content of a file.txt that have numbers s
func GetDatas(name_file string) []int {
	d, _ := os.ReadFile(name_file)
	content := string(d)
	var slice []int
	save := ""
	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			slice = append(slice, Atoi(string(save)))
			save = ""
		} else if content[i] != '\r' {
			save += string(content[i])
		}
	}

	return slice
}

func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// Sort an int slice
func Sort(t []int) {
	for i := 0; i < len(t)-1; i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i] > t[j] {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
}
