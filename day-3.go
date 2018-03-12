package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	inputNumber, _ := strconv.Atoi(readInputLine())
	if inputNumber%2 == 1 {
		fmt.Println("Weird")
	} else if inputNumber > 20 {
		fmt.Println("Not Weird")
	} else if inputNumber > 5 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}

}

func readInputLine() string {
	text, _ := reader.ReadString('\n')
	return text
}
