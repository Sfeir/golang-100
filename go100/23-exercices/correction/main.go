package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
}

func (p *Person) Designation() string {
	if p == nil {
		return "Anonyme"
	}
	return p.Firstname + " " + p.Lastname
}

func main() {
	person := Person{"Jack", "Sparo"}
	fmt.Println(person.Designation())
}
