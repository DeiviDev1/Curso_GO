package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"Jose Joao":      1325.25,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}
	//adiciona
	funcsESalarios["Rafael Filho"] = 1350.0
	funcsESalarios["Deivi Roch"] = 100000.0
	//deleta,
	delete(funcsESalarios, "Inexistente")
	//para puxar apenas um elemento deve-se usar "_ traço baixo "
	//exemplo  for _, salario := range.. desta forma mostrará apenas o salario
	for nome, salario := range funcsESalarios {

		fmt.Println(nome, salario)
	}
}
