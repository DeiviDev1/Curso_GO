package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qntde int) {

	for i := 0; i < qntde; i++ {

		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)

	}
}

func main() {
	//fale("Maria", "pq vc não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc !", 1)

	//go fale("maria", "ei...",500)
	//go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabens!", 5)
}
