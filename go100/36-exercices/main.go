package main

import (
	"fmt"
)

//TODO: modifiez la fonction fibonacci de l'exercice précédent pour qu'elle envoie les valeurs sur le channel c jusqu'à ce qu'elle reçoive
//un signal d'arrêt.
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
