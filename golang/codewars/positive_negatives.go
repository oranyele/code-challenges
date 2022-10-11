package main

import (
	"fmt"
)

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	numPositives := 0
	sumNegatives := 0

	for _, v := range numbers {
		if v > 0 {
			numPositives++
		} else if v < 0 {
			sumNegatives += v
		}
	}
	res = append(res, numPositives)
	res = append(res, sumNegatives)
	return res // your code here
}

func main() {
	fmt.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}
