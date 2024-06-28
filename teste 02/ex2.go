package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func troca(vetor []int, i, j int) {
	vetor[i], vetor[j] = vetor[j], vetor[i]
}


func trocaOpostosSeMenor(vetor []int, n int) {
	for i := 0; i < n/2; i++ {
		oposto := n - 1 - i
		if vetor[i] < vetor[oposto] {
			troca(vetor, i, oposto)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)


	scanner.Scan()
	numCasos, _ := strconv.Atoi(scanner.Text())

	for caso := 0; caso < numCasos; caso++ {

		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())


		scanner.Scan()
		vetorStr := strings.Fields(scanner.Text())
		vetor := make([]int, n)
		for i, v := range vetorStr {
			vetor[i], _ = strconv.Atoi(v)
		}


		trocaOpostosSeMenor(vetor, n)


		for i, v := range vetor {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro na leitura da entrada:", err)
	}
}
