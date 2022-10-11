package main

import "fmt"

func Rps(p1, p2 string) string {
	if (p1 == "scissors") && (p2 == "paper") {
		return "Player 1 won!"
	} else if (p1 == "paper") && (p2 == "rock") {
		return "Player 1 won!"
	} else if (p1 == "rock") && (p2 == "scissors") {
		return "Player 1 won!"
	} else if (p2 == "scissors") && (p1 == "paper") {
		return "Player 2 won!"
	} else if (p2 == "paper") && (p1 == "rock") {
		return "Player 2 won!"
	} else if (p2 == "rock") && (p1 == "scissors") {
		return "Player 2 won!"
	}
	return "Draw!"
}

func main() {
	fmt.Println(Rps("scissors", "paper"))
}

// Best solution
// var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

// func Rps(a, b string) string {
//   if a == b {
//     return "Draw!"
//   }
//   if m[a] == b {
//     return "Player 2 won!"
//   }
//   return "Player 1 won!"
// }
