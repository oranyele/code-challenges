package main

import "fmt"

func ExpressionMatter(a int, b int, c int) int {
	// your code here
	var maiorValor = 0
	var temp = 0

	temp = a + b + c
	maiorValor = qualOMaior(maiorValor, temp)

	temp = a * b * c
	maiorValor = qualOMaior(maiorValor, temp)

	temp = (a + b) * c
	maiorValor = qualOMaior(maiorValor, temp)

	temp = a * (b + c)
	maiorValor = qualOMaior(maiorValor, temp)

	temp = a + b*c
	maiorValor = qualOMaior(maiorValor, temp)
	return maiorValor
}

func qualOMaior(old int, temp int) int {
	fmt.Printf("%d - %d\n", old, temp)
	if temp > old {
		return temp
	} else {
		return old
	}
}

func main() {
	fmt.Println(ExpressionMatter(2, 1, 2))
}
