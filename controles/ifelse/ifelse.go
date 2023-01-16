package main

import "fmt"

func imprimirResultado(nota float64) {
	// não colocar paranteses na definição if em go , DEVEM ESTAR EM BLOCO
	if nota >= 7 {
		fmt.Println("Aprovado com nota ", nota)
	} else {
		fmt.Println("Reprovado com a nota ", nota)
	}
}

func main() {

	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
