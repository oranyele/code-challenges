package main

import "fmt"

func DNAtoRNA(dna string) string {
	// your code here
	rna := ""
	for i := 0; i < len(dna); i++ {
		if string(dna[i]) == "T" {
			rna += "U"
		} else {
			rna += string(dna[i])
		}

	}
	return rna
}

func main() {
	fmt.Println(DNAtoRNA("GCAT"))
}
