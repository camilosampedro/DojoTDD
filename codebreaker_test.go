package codebreaker

import "testing"

func TestTodoBien(t *testing.T) {
  codebreaker := CodeBreaker{"1839"}
	expected := "XXXX"
	actual := evaluar(codebreaker,"1839")
	if actual != expected {
		t.Errorf("Test falló. Esperado %s, actual: %s", expected, actual)
	}
}

func TestDosX(t *testing.T) {
  codebreaker := CodeBreaker{"1839"}
	expected := "XX"
	actual := evaluar(codebreaker,"1826")
	if actual != expected {
		t.Errorf("Test falló. Esperado %s, actual: %s", expected, actual)
	}
}

func TestTresX(t *testing.T) {
  codebreaker := CodeBreaker{"1839"}
	expected := "XXX"
	actual := evaluar(codebreaker,"1836")
	if actual != expected {
		t.Errorf("Test falló. Esperado %s, actual: %s", expected, actual)
	}
}

func TestTres_(t *testing.T) {
  codebreaker := CodeBreaker{"1839"}
	expected := "___"
	actual := evaluar(codebreaker,"3186")
	if actual != expected {
		t.Errorf("Test falló. Esperado %s, actual: %s", expected, actual)
	}
}
