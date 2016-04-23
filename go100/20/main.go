package main

import (
	"fmt"
)

func fonction(prefix string, b ...string) {
	fmt.Printf("b est égale à %v et est du type %T\n", b, b)
	for _, v := range b {
		fmt.Printf("%s%s\n", prefix, v)
	}
}

func main() {
	fonction(">>> ", "a", "b", "c", "hello")
}
