package main

import "fmt"

type person struct {
	firstName    string
	lastName     string
	favCountries []string
}

func main() {

	//p1 := person{firstName: "Tharun", lastName: "Rede"}
	//p1.favCountries = []string{"India", "UK", "USA"}
	p1 := person{"Tharun", "Rede", []string{"India", "UK", "USA"}}
	p2 := person{"Adam", "Robert", []string{"Canada", "Netherlands", "Finland"}}

	fmt.Println("First person struct: ", p1)
	fmt.Println("Person First name: ", p1.firstName)

	p1.favCountries = append(p1.favCountries, "Africa")
	fmt.Println("First person Fav Countries are: ")

	for index, fav := range p1.favCountries {
		fmt.Println(index, fav)

	}

	fmt.Println("Second person struct: ", p2)
	fmt.Println("Person First name: ", p2.firstName)
	fmt.Println("First person Fav Countries are: ")

	for index, fav := range p2.favCountries {
		fmt.Println(index, fav)

	}

	fmt.Println("Problem-2")
	personMap := make(map[string]person)

	personMap[p1.lastName] = p1
	personMap[p2.lastName] = p2

	fmt.Println(personMap)

	for key, _ := range personMap {
		fmt.Println(personMap[key].favCountries)
		for k1, f1 := range personMap[key].favCountries {
			fmt.Println(k1, f1)

		}
	}

}
