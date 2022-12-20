package main

import (
	"fmt"

	"assignments/arithmetic"
)

func main() {

	var a, b int
	fmt.Println("Enter a Value: ")
	fmt.Scanln(&a)
	fmt.Println("Enter b Value: 10")
	fmt.Scanln(&b)

	fmt.Printf("Sum of %v and %v is : %v\n", a, b, arithmetic.Sum(a, b))
	fmt.Printf("Difference of %v and %v is : %v\n", a, b, arithmetic.Difference(a, b))
	fmt.Printf("Product of %v and %v is : %v\n", a, b, arithmetic.Product(a, b))

	quotient, remainder := arithmetic.Division(a, b)

	fmt.Printf("Quotient of %v divided by %v is : %v\n", a, b, quotient)
	fmt.Printf("Remainder of %v divided by %v is : %v\n", a, b, remainder)

}
