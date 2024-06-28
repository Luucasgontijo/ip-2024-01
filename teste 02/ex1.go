package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func inverterPalavra(palavra string) string {
	runes := []rune(palavra)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func inverterFrase(frase string) string {
	palavras := strings.Fields(frase)
	for i, palavra := range palavras {
		palavras[i] = inverterPalavra(palavra)
	}
	return strings.Join(palavras, " ")
}

func main() {
	fmt.Println("insira as palavras")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linha := scanner.Text()
		if linha != "" {
			fraseInvertida := inverterFrase(linha)
			fmt.Println(fraseInvertida)
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro na leitura da entrada:", err)
	}
}
