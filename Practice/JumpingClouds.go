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
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32, n int32) int32 {
    // Write your code here
	var jumps int32
	jumps = 0
	l := int(n)

	for i := 0; i < l ;{

		jumps++

		if i == l-3 || i == l-2{
			break
		}

		if c[i+2] == 0 {
			i = i + 2
		}else{
			i = i + 1
		}
	}

	return jumps
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var c []int32

    for i := 0; i < int(n); i++ {
        cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
        checkError(err)
        cItem := int32(cItemTemp)
        c = append(c, cItem)
    }

    result := jumpingOnClouds(c,n)

    fmt.Fprintf(writer, "%d\n", result)

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
