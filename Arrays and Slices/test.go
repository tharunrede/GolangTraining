package main

import "fmt"

func main() {

	x := [4]int{1, 2, 3}
	fmt.Println(x)

	x = []int{1, 2}
	x = x[2:4]
}
