package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println()
	fmt.Println("Problem-5")

	fmt.Println("Enter your String(Golang Training/ other): \n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_str := scanner.Text()

	fmt.Println("Input value is: ", input_str)

	fmt.Printf("type %T and value %v \n", input_str, input_str)

	if input_str == "Golang Training" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}

}
