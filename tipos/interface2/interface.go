package main

import "fmt"

type esportivo interface {
	LigarTurbo()
}
type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) LigarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.LigarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.LigarTurbo()

	fmt.Println(carro1, carro2)
}
