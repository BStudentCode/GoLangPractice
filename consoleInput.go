package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What year were you born?: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //, _ means if errors then ignore and continue program, strconv contains string methods
	fmt.Printf("You will be %d years old at the end of 2021", 2021-input)
}
