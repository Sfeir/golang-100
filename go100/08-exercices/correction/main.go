package main

import (
	"errors"
	"fmt"
)

//inverse prend un nombre décimal en paramètre et retourne son inverse. Si le nombre est nul, une erreur est retournée.
func inverse(x float64) (float64, error) {
	if x == 0 {
		return 0, errors.New("L'inverse de zéro n'existe pas.")
	}
	return 1 / x, nil
}

func main() {
	// Calcul de l'inverse de 3
	if inv, err := inverse(3); err == nil {
		fmt.Printf("L'inverse de %f est %f\n", 3.0, inv)
	} else {
		fmt.Println(err)
	}

	// Calcul de l'inverse de 0 !
	if inv, err := inverse(0); err == nil {
		fmt.Printf("L'inverse de %f est %f\n", 0.0, inv)
	} else {
		fmt.Println(err)
	}
}
