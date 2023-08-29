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
 * Complete the 'utopianTree' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */
// cycle   height    2^n
// 0       1         1
// 1       2         2
// 2       3         4
// 3       6         8
// 4       7         16
// 5       14        32
// 6       15        64
// 7       30        128
// 8       31        256
// 9       62        512
// 10      63        1024
// 11      126       2048
// 12      127       4096
//
// oddCycle: 2^((cycle+1)/2 + 1) - 2
// evenCycle: 2^(cycle/2 + 1) - 1
func utopianTree(n int32) int32 {
	if n%2 == 1 {
		return (1 << ((n+1)/2 + 1)) - 2
	}
	return (1 << (n/2 + 1)) - 1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		result := utopianTree(n)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
