package main

import "testing"

// Test 1: PASA - Comprueba el cálculo exacto de la letra de un DNI
func TestCalcularLetraDNI(t *testing.T) {
	numero := 12345678 // Para este número la letra oficial es Z
	letraEsperada := "Z"
	letraObtenida := CalcularLetraDNI(numero)

	if letraObtenida != letraEsperada {
		t.Errorf("Error: se esperaba la letra %s pero se obtuvo %s", letraEsperada, letraObtenida)
	}
}

// Test 2: PASA - Comprueba la validación de un DNI completo correcto
func TestValidarDNICorrecto(t *testing.T) {
	dniValido := "12345678Z"
	if !ValidarDNI(dniValido) {
		t.Errorf("Error: el DNI %s debería ser válido", dniValido)
	}
}

// Test 3: FALLA INTENCIONADAMENTE - Afirma que un DNI con letra errónea es válido
func TestValidarDNIFallido(t *testing.T) {
	dniFalso := "12345678A" // La letra correcta para 12345678 es Z, no A

	// Afirmamos que la función debe devolver true para forzar el FALLO
	if !ValidarDNI(dniFalso) {
		t.Errorf("Fallo intencionado: Se esperaba que el DNI %s fuera válido pero dio falso", dniFalso)
	}
}
