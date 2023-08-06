package main

import "fmt"

func main() {
	person := Person{
		Name:  "Mohammed Sadiq",
		Age:   24,
		Email: "mohammed_sadiq@thbs.com",
	}

	fmt.Printf(" Name : %s\n Age : %d\n Email : %s", person.Name, person.Age, person.Email)
}
