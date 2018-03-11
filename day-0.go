// https://www.hackerrank.com/challenges/30-hello-world/problem

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Print("Hello, World.\n")
	fmt.Println(text)
}
