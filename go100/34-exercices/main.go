package main

import (
	"fmt"
)

//TODO: implémentez la fonction fibonacci qui devra envoyer les valeur sur le channel c.
func fibonacci(n int, c chan int) {
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		fmt.Println(i)
	}
}
