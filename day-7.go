// https://www.hackerrank.com/challenges/30-arrays/problem

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
	arrayLength, _ := strconv.Atoi(readInputLine())
	valueArrayString := strings.Split(readInputLine(), " ")
	var result string
	for i := arrayLength - 1; i > -1; i-- {
		result += (valueArrayString[i] + " ")
	}
	fmt.Println(result)
}

func readInputLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
