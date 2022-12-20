package arithmetic

func Sum(a int, b int) int {
	return a + b
}

func Difference(a, b int) int {
	return a - b
}

func Product(a, b int) int {
	return a * b
}

func Division(a, b int) (int, int) {
	return a / b, a % b
}
