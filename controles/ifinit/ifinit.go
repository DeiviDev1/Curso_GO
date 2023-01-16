package main

import (
	"fmt"
	"math/rand"
	"time"
)

// gera numero aleatorio usando data e segundos
func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {

	if i := numeroAleatorio(); i > 5 { // tbm é suportado com switch
		fmt.Println("Ganhou !!!!")
	} else {
		fmt.Println("Perdeu...!!!")
	}
}
