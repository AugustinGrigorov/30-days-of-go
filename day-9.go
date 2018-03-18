// https://www.hackerrank.com/challenges/30-recursion/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	n, _ := strconv.Atoi(readInputLine())
	fmt.Println(factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func readInputLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
