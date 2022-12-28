package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Problem-1")

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Println("Problem-2")

	i := 1
	for i < 50 {

		if i%2 == 1 {
			fmt.Println(i)
		}
		i++
	}
	fmt.Println()
	fmt.Println("Problem-3")

	j := 1
	for {

		if j%2 == 0 {
			fmt.Println(j)

		} else if j > 50 {
			break
		}
		j++
	}

	fmt.Println()
	fmt.Println("Problem-4")

	for i = 50; i <= 105; i++ {
		fmt.Println(i / 6)
	}

	fmt.Println()
	fmt.Println("Problem-5")

	fmt.Print("Enter your String(Golang Tutorial/ other): \n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_str := scanner.Text()
	fmt.Println("Input value is: ", input_str)

	//fmt.Printf("type %T and value %v \n", input_str, input_str)

	if input_str == "Golang Training" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}

	fmt.Println()
	fmt.Println("Problem-6")

	for k := 1; k <= 100; k++ {

		if k%2 == 0 && k%4 == 0 {
			fmt.Println("Golang tutorial")
		} else if k%4 == 0 {
			fmt.Println("tutorial")

		} else if k%2 == 0 {
			fmt.Println("Golang")
		} else {
			fmt.Println(k)
		}
	}

}
