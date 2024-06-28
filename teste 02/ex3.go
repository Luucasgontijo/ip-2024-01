package main

import (
	"fmt"
	"sort"
)

func ordenarVetor(vetor []int) {
	sort.Ints(vetor)
}

func main() {

	vetor := []int{5, 2, 9, 1, 5, 6}

	fmt.Println("Vetor original:", vetor)


	ordenarVetor(vetor)

	fmt.Println("Vetor ordenado:", vetor)
}
