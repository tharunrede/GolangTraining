package main

import (
	"fmt"
)

func main() {

	var arr [5]int
	arr[0] = 10
	arr[1] = 65
	arr[2] = 71
	arr[3] = 1
	arr[4] = 6
	fmt.Println("Problem 1: ")
	fmt.Printf("The variable arr is of type: %T\n", arr)

	var slice_int = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Problem 2: ")
	fmt.Printf("The variable slice_int is of type: %T\n", slice_int)

	fmt.Println("Problem 3: ")

	fmt.Println("slice1 is: ", slice_int[:5])

	fmt.Println("slice2 is: ", slice_int[5:])

	fmt.Println("slice3 is: ", slice_int[2:7])

	fmt.Println("slice4 is: ", slice_int[3:8])

	fmt.Println("Problem 4: ")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("x array is: ", x)

	x = append(x, 52)
	fmt.Println("x after adding 52 is: ", x)

	x = append(x, 53, 54, 55)
	fmt.Println("x after adding 53, 54, 55 is: ", x)
	y := []int{56, 57, 58, 59, 60}
	fmt.Println("y array is: ", y)
	x = append(x, y...)

	fmt.Println("x after appending array y is: ", x)

	fmt.Println("Problem 5: ")

}
