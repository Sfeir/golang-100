package main

import "fmt"

func realSyracuse(val, currentItern, maxIter int) (bool, int) {
	if val == 1 {
		return true, currentItern
	} else if currentItern >= maxIter {
		return false, currentItern
	} else if val % 2 == 0 {
		return realSyracuse(val / 2, currentItern+1, maxIter)
	} else {
		return realSyracuse(3*val + 1, currentItern+1, maxIter)
	}
}

func recursiveSyracuse(val, nbMaxIteration int) (bool, int) {
	return realSyracuse(val, 0, nbMaxIteration)
}

//syracuse prend en paramètre un entier et un nombre max d'itérations à effectuer.
//Elle retourne un booléen indiquant si le nombre 1 a été atteint, suivi du nombre d'itérations qui ont été effectuées.
func syracuse(start int, maxIterations int) (bool, int) {
	//cas start < 1
	if start < 1 {
		return false, 0
	}
	//algorithme
	i := 0
	for n := start; n > 1 && i < maxIterations; i++ {
		fmt.Printf("iteration %d, n = %d\n", i, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return i < maxIterations, i
}

func main() {
	const maxIterationCount = 1000
	const integer = 98375
	reached, iterationCount := syracuse(integer, maxIterationCount)
	if reached {
		fmt.Printf("L'algorithme a permis d'atteindre le nombre 1 à partir du nombre %d en %d itérations\n", integer, iterationCount)
	} else {
		fmt.Printf("L'algorithme n'a pas permis d'atteindre le nombre 1 à partir du nombre %d en moins de %d itérations\n", integer, maxIterationCount)
	}
}
