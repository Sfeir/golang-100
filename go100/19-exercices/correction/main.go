package main

import "fmt"

// fibonacci qui retourne une fonction (plus précisément, une closure) qui retourne un entier.
func fibonacci() func() int {
	prev, n := 0, 1
	return func() int {
		// le code ci-dessous est équivalent à:
		// oldN = n
		// n = prev + n
		// prev = oldN
		prev, n = n, prev+n
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
