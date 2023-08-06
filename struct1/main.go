package main

import "fmt"

func main() {

	person1 := Person{
		Name:  "Mohammed Sadiq",
		Age:   23,
		Email: "mohammed_sadiq@thbs.com",
	}

	person2 := Person{
		Name:  "Sinchana Gopi",
		Age:   23,
		Email: "sinchana_gopi@thbs.com",
	}

	fmt.Printf(" Name : %s\n Age : %d\n Email : %s\n\n", person1.Name, person1.Age, person1.Email)
	fmt.Printf(" Name : %s\n Age : %d\n Email : %s", person2.Name, person2.Age, person2.Email)
}