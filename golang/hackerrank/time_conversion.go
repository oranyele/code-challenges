package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	horas := strings.Split(s, ":")
	ampm := horas[2][2:4]
	fmt.Println(ampm)
	formatada := ""
	if (ampm == "PM") && (horas[0] == "12") {
		formatada += "00"
	} else {
		novaHora, _ := strconv.Atoi(horas[0])
		novaHora += 12
		formatada = strconv.Itoa(novaHora)
	}
	formatada += fmt.Sprintf(":%s:%s", horas[1], horas[2][0:2])
	return formatada
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// s := readLine(reader)
	s := "00:05:45PM"
	result := timeConversion(s)

	fmt.Printf("%s\n", result)

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
