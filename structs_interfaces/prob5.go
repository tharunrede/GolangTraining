package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter your Sentence: \n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_str := scanner.Text()
	//fmt.Println("Input value is: ", input_str)

	wordsmap := map[string]int{}

	//fmt.Printf("total of words:%v \n", strings.Split(input_str, " "))

	words := strings.Fields(input_str)
	keys := []string{}

	for i := 0; i < len(words); i++ {

		ind_word := words[i]
		//fmt.Println(len(ind_word))

		found := false

		// iterate over the slice
		for _, v := range keys {

			// check if the strings match
			if v == ind_word {
				found = true
				//fmt.Println("The slice contains", ind_word, "at index")

				break
			}
		}

		//fmt.Println("keys:", keys)

		// if string not found
		if found == false {
			wordsmap[ind_word] = 1
			keys = append(keys, ind_word)
			//fmt.Println("here found false")
		} else {
			wordsmap[ind_word] += 1
		}

	}
	//fmt.Println("keys out:", keys)
	fmt.Println("Final Map: ", wordsmap)
}
