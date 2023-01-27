package main

import "fmt"

func main() {
	// map dentro de map
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Siilva": 15465.78,
			"Guga Pereira":    8456.78,
		},
		"J": {
			"Jose Joao":  11235.45,
			"Joana Dark": 20153.13,
		},

		"D": {
			"Deiv Roch":      25145.15,
			"Daniel Pereira": 5225.15,
		},
	}
	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
