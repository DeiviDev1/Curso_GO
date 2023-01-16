package main

import "fmt"

// Go não tem operador ternário
func obeterResultado(nota float64) string {

	if nota > 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {

	fmt.Println(obeterResultado(6.2))
}
