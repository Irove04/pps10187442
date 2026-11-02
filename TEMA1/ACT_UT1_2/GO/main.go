package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Tabla oficial de correspondencia de letras del DNI
const letrasDNI = "TRWAGMYFPDXBNJZSQVHLCKE"

// CalcularLetraDNI devuelve la letra oficial según los números del DNI
func CalcularLetraDNI(numero int) string {
	indice := numero % 23
	return string(letrasDNI[indice])
}

// ValidarDNI comprueba si un DNI (ej: "12345678Z") es correcto
func ValidarDNI(dni string) bool {
	dni = strings.TrimSpace(strings.ToUpper(dni))
	if len(dni) != 9 {
		return false
	}

	numeroStr := dni[:8]
	letra := string(dni[8])

	numero, err := strconv.Atoi(numeroStr)
	if err != nil {
		return false
	}

	return CalcularLetraDNI(numero) == letra
}

func main() {
	numeroDNI := 12345678
	letra := CalcularLetraDNI(numeroDNI)
	fmt.Printf("Para el número %d, la letra calculada es: %s\n", numeroDNI, letra)

	dniCompleto := "12345678Z"
	fmt.Printf("¿Es válido el DNI %s?: %t\n", dniCompleto, ValidarDNI(dniCompleto))
}
