package main

import "fmt"

//TODO: créez un type Person qui contiendra le prénom et le nom d'une personne.
//type Person...

//TODO: créez une méthode 'Designation' sur le type *Person qui retourne la concaténation du prénom et du nom.
//La méthode devra gérer le cas où la personne est nil.
//func...

func main() {
	person := Person{"Jack", "Sparo"}
	fmt.Println(person.Designation())
}
