package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

//Déclarer la méthode M sur le type *T va faire que *T implémente automatiquement l'interface I.
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

//Déclarer la méthode M sur le type F va faire que F implémente automatiquement l'interface I.
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
