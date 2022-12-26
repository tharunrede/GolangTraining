package main

import "fmt"

func main() {

	states := [4]string{"Arizona", "New Mexico", "New York", "California"}

	fmt.Println("States before changing: ", states)
	updateThirdElement(&states[2])
	fmt.Println("States after changing: ", states)

}

func updateThirdElement(states *string) {
	*states = "Texas"
	fmt.Println(*states)

}

// func main() {

// 	states := [4]string{"Arizona", "New Mexico", "New York", "California"}

// 	fmt.Println("States before changing: ", states)
// 	updateThirdElement(&states)
// 	fmt.Println("States after changing: ", states)

// }

// func updateThirdElement(states *[4]string) {
// 	(*states)[2] = "Texas"
// 	fmt.Println(*states)

// }
