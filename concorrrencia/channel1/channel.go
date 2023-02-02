package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "teste" //enviando dados para um canal(escrita)
	<-ch          //recebebdi dadis do canal (leitura)

	ch <- "teste 2"
	fmt.Println(<-ch)

}
