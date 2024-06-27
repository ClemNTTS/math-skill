package main

import (
	"fmt"
	"math"
	"os"

	"math-skills/lib"
)

func main() {
	//Protection for wrong terminal input
	if len(os.Args) != 2 {
		fmt.Println("Error: to much or missing parameters")
		return
	}
	//Get datas in a int slice and sort it
	name_file := os.Args[1]
	tab := lib.GetDatas(name_file)
	lib.Sort(tab)

	//Apply and print all functions and rounded results
	fmt.Printf("Average : %d\n", int(math.Round(lib.Average(tab))))
	fmt.Printf("Median : %d\n", int(math.Round(float64(lib.Median(tab)))))
	fmt.Printf("Variation : %d\n", int(math.Round(float64(lib.Variance(int(lib.Average(tab)), tab)))))
	fmt.Printf("Standard Deviation : %d\n", lib.StandardDeviation(float64(lib.Variance(int(lib.Average(tab)), tab))))

}
