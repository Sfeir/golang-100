package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Le type T implémente I (au choix)
func (t T) M() {
	fmt.Println(t.S)
}

// Le type *T implémente I (au choix)
func (t *T) M() {
	fmt.Println(t.S)
}

func main() {
	// Si le type T implémente I, on peut faire
	var i I = T{"hello"}
	i.M()

	// Mais si le type *T implémente I, on peut faire
	var i I = &T{"hello"}
	i.M()
}
