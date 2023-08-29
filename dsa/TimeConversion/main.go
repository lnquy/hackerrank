package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	sb := []byte(s)
	if s[len(s)-2:] == "AM" {
		sb = sb[:len(s)-2]
		if s[:2] == "12" {
			sb[0] = '0'
			sb[1] = '0'
		}
		return string(sb)
	}

	// PM
	sb = sb[:len(s)-2]
	if s[:2] == "12" {
		return string(sb)
	}
	hour, _ := strconv.Atoi(s[:2])
	hour += 12
	sb[0] = strconv.Itoa(hour / 10)[0]
	sb[1] = strconv.Itoa(hour % 10)[0]
	return string(sb)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
