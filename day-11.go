// https://www.hackerrank.com/challenges/30-2d-arrays/problem

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]float64
	for scanner.Scan() {
		grid = append(grid, convertInputToNumericArray(scanner.Text()))
	}
	largestHourglassSum := math.Inf(-1)
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid)-2; j++ {
			var hourglassSum float64
			hourglassSum += grid[i][j]
			hourglassSum += grid[i][j+1]
			hourglassSum += grid[i][j+2]
			hourglassSum += grid[i+1][j+1]
			hourglassSum += grid[i+2][j]
			hourglassSum += grid[i+2][j+1]
			hourglassSum += grid[i+2][j+2]
			if hourglassSum > largestHourglassSum {
				largestHourglassSum = hourglassSum
			}
		}
	}
	fmt.Println(largestHourglassSum)
}

func convertInputToNumericArray(input string) []float64 {
	inputStringArray := strings.Split(input, " ")
	result := make([]float64, len(inputStringArray))
	for index, inputString := range inputStringArray {
		result[index], _ = strconv.ParseFloat(inputString, 64)
	}
	return result
}
