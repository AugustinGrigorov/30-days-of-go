// https://www.hackerrank.com/challenges/30-operators/problem

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
	mealCost, _ := strconv.ParseFloat(readInputLine(), 64)
	tipPercent, _ := strconv.ParseFloat(readInputLine(), 64)
	taxPercent, _ := strconv.ParseFloat(readInputLine(), 64)
	tipAmount := (mealCost * tipPercent) / 100
	taxAmount := (mealCost * taxPercent) / 100
	totalRoundedCost := int(mealCost + tipAmount + taxAmount + 0.5) // The 0.5 is a hack around the absence of math.Round in outdated versions of Go
	fmt.Printf("The total meal cost is %v dollars.", totalRoundedCost)
}

func readInputLine() string {
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}
