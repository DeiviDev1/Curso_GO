package matematica

import "testing"

const erroPadrao = " Valor esperadeo %v, mas o resultado encontrado foi %v."

func TestMed(t *testing.T) {
	valotEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valotEsperado {
		t.Errorf(erroPadrao, valotEsperado, valor)
	}
}
