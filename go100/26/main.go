package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("M() says: <nil>")
		return
	}
	fmt.Println("M() says: " + t.S)
}

func main() {
	var i I

	fmt.Println("Cas 1 :")
	i = &T{"hello"}
	describe(i)
	i.M()

	fmt.Println("\nCas 2 :")
	var t *T
	// ici, t = nil
	i = t
	describe(i)
	i.M()

	fmt.Println("\nCas 3 :")
	i = nil
	describe(i)
	i.M() // panic!
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
