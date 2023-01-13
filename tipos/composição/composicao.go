package main

import "fmt"

// interface composta
type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazeBaliza()
}
type exclusivo interface {
	painelSolar()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	exclusivo
	//pode adicionar mais métodos

}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("TURBO...")
}
func (b bmw7) fazeBaliza() {
	fmt.Println("Baliza..")
}
func (bmw7) painelSolar() {
	fmt.Println("carregar solar")
}
func main() {
	var b esportivoLuxuoso = bmw7{}
	b.fazeBaliza()
	b.ligarTurbo()
	b.painelSolar()

}
