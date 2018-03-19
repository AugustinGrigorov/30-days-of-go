// https://www.hackerrank.com/challenges/30-binary-numbers/problem

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
	var binaryRepresentation []int
	var currentStreak int
	var longestStreak int
	numberToConvert, _ := strconv.Atoi(readInputLine())
	for numberToConvert > 0 {
		binaryRepresentation = append(binaryRepresentation, numberToConvert%2)
		numberToConvert = numberToConvert / 2
	}
	for _, binaryNumber := range binaryRepresentation {
		if binaryNumber == 1 {
			currentStreak++
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}

	}
	fmt.Println(longestStreak)
}

func readInputLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
