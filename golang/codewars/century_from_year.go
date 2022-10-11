package main

import "fmt"

func century(year int) int {
	// Finish this :)
	if year%100 > 0 {
		return (year / 100) + 1
	}
	return year / 100
}

func main() {
	fmt.Println(century(1705))
	fmt.Println(century(1900))
	fmt.Println(century(1601))
	fmt.Println(century(2000))
}
