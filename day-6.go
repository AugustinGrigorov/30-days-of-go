// https://www.hackerrank.com/challenges/30-review-loop/problem

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	for scanner.Scan() {
		var evenString string
		var oddString string
		for index, char := range scanner.Text() {
			addition := string(char)
			if index%2 == 0 {
				evenString += addition
			} else {
				oddString += addition
			}
		}
		fmt.Printf("%v %v\n", evenString, oddString)
	}
}
