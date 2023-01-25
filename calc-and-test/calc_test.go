package main

import "testing"

func TestSomar(t *testing.T) {
	teste := somar(2.5, 2.5)
	resultado := 5.0

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor  retornado", teste)
	}
}

func TestSubtrair(t *testing.T) {
	teste := subtrair(7.5, 1, 1.5)
	resultado := 5.0

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor  retornado", teste)
	}
}

func TestMultiplicar(t *testing.T) {
	teste := multiplicar(5, 0.5, 2)
	resultado := 5.0

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor  retornado", teste)
	}
}

func TestDividir(t *testing.T) {
	teste := dividir(100, 8, 2.5)
	resultado := 5.0

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor  retornado", teste)
	}
}
