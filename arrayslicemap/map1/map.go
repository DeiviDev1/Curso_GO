package main

import "fmt"

func main() {
	// maps devem ser inicializados (make)
	//var aprovados map[int]string

	aprovados := make(map[int]string)
	//chave valor
	aprovados[12345678] = "Maria"
	aprovados[28244825] = "Pedro"
	aprovados[35846712] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	//lê o map
	fmt.Println(aprovados[12345678])

	//deleta um elemento do map
	delete(aprovados, 12345678)
	fmt.Println(aprovados[12345678])
}
