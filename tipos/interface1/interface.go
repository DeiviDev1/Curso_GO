package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {

	return p.nome + " " + p.sobrenome
}
func (p produto) toString() string {

	return fmt.Sprintf("%s - R$ %.3f", p.nome, p.preco)
}

func imprimir(x imprimivel) {

	fmt.Println(x.toString())
}

// a variavel "coisa" assume multiplas formas
func main() {
	var coisa imprimivel = pessoa{"Deivion", "Rocha"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"calça", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	var coisa1 imprimivel = produto{"Pc", 1.90000}
	fmt.Println(coisa1.toString())
	imprimir(coisa1)

}
