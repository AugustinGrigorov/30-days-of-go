// https://www.hackerrank.com/challenges/30-data-types/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var input []string

	// Declare second integer, double, and String variables.
	var intVar uint64
	var doubleVar float64
	var stringVar string

	// Read and save an integer, double, and String to your variables.
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	intText := input[0]
	intVar, _ = strconv.ParseUint(intText, 10, 64)
	doubleText := input[1]
	doubleVar, _ = strconv.ParseFloat(doubleText, 64)
	stringVar = input[2]

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + intVar)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+doubleVar)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + stringVar)
}
