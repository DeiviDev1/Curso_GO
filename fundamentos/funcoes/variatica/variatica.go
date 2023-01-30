package main

//import "fmt"

// função variatica represetada pelos( ...) pode ser inseridos varios parametros coforme linha15 do código
func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

//deixar essa função desativada( não pode existir 2 funçoes maain na msm pasta)
//func main() {

//fmt.Printf("Média: %2.f", media(7.7, 8.1, 8.3, 9.0, 5.9))
//}
