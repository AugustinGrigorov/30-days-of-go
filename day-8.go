// https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	phonebook := make(map[string]string)
	scanner.Scan()
	numberOfRecords, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < numberOfRecords; i++ {
		scanner.Scan()
		record := strings.Split(scanner.Text(), " ")
		phonebook[record[0]] = record[1]
	}

	for scanner.Scan() {
		name := scanner.Text()
		if phonebook[name] != "" {
			fmt.Printf("%v=%v\n", name, phonebook[name])
		} else {
			fmt.Println("Not found")
		}
	}
}
