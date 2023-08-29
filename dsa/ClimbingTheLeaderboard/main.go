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
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func climbingLeaderboard(ranked []int32, player []int32) (result []int32) {
	for i := 0; i < len(ranked)-1; i++ {
		if ranked[i] == ranked[i+1] {
			ranked = append(ranked[:i], ranked[i+1:]...) // De-dup
			i--
		}
	}

	for _, point := range player {
		result = append(result, int32(binarySearch(ranked, point)))
	}
	return result
}

func binarySearch(a []int32, v int32) int {
	if v >= a[0] {
		return 1
	}
	if v == a[len(a)-1] {
		return len(a)
	}
	if v < a[len(a)-1] {
		return len(a) + 1
	}

	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] == v {
			return mid + 1
		}
		if a[mid] < v {
			if mid > 0 && a[mid-1] > v {
				return mid + 1
			}
			r = mid - 1
			continue
		}

		if a[mid] > v {
			if mid < len(a)-1 && a[mid+1] < v {
				return mid + 1 + 1
			}
			l = mid + 1
			continue
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
