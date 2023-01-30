package main

import "fmt"

// função vazia, não recebe e n retorna
func f1() {

	fmt.Println("F1")

}

// função com parametros não retorna nada.
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)

}

// função com retorno
func f3() string {

	return "F3"
}

// recebe paramatros e retorna
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)

}

// retornar multiplos valores
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"

}

func main() {

	f1()
	f2("Parametro1", "Parametro2")

	r3, r4 := f3(), f4("Paramatro1", "Paramtro2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
