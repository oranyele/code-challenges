package main

import (
	"strings"
)

func Points(games []string) int {
	// your code here!
	totalPoints := 0
	for _, v := range games {
		points := strings.Split(v, ":")
		if points[0] > points[1] {
			totalPoints += 3
		} else if points[0] == points[1] {
			totalPoints += 1
		}
	}
	return totalPoints
}

// if x > y: 3 points (win)
// if x < y: 0 points (loss)
// if x = y: 1 point (tie)

func main() {
	games := []string{
		"3:1",
		"2:2",
		"0:1",
	}
	Points(games)
}
