package main

import (
	"fmt"
	"time"
)

// channel é um tipo

func canal(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base

}

func main() {

	c := make(chan int)
	go canal(2, c)

	x, y := <-c, <-c // recebendo dados do canal
	fmt.Println(x, y)
	fmt.Println(<-c)

}
