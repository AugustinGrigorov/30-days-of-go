// https://www.hackerrank.com/challenges/30-loops/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	n, _ := strconv.Atoi(readInputLine())
	for i := 1; i < 11; i++ {
		fmt.Printf("%v x %v = %v\n", n, i, n*i)
	}
}

func readInputLine() string {
	text, _ := reader.ReadString('\n')
	return text
}
