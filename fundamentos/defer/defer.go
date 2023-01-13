package main

import "fmt"

func obterValorAprovado(num int) int {
	defer fmt.Println("fim!")
	if num > 5000 {
		fmt.Println("Valor alto..")
		return 5000
	} else {
		fmt.Println("valor menor...")
		return num
	}

}
func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))

}
