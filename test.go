package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your String: \n")

	input, _ := consoleReader.ReadString('\n')

	fmt.Println("input is ", input)

	flag := strings.Compare(input, "Golang tutorial")

	//fmt.Printf("type %T and value is %v\n", flag, flag)

	if flag == 1 {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}

}
