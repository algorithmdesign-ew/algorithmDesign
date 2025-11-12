package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var opcion int
	var texto string

	for {
		fmt.Println("\n=== ALGORITMO DE ENCRIPTACIÓN ===")
		fmt.Println("1. Cifrar")
		fmt.Println("2. Descifrar")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Print("\nIngrese el texto a cifrar: ")
			fmt.Scanln(&texto)
			resultado := cifrar(texto)
			fmt.Printf("\n✓ Texto cifrado: %s\n", resultado)

		case 2:
			fmt.Print("\nIngrese el texto cifrado: ")
			fmt.Scanln(&texto)
			resultado := descifrar(texto)
			fmt.Printf("\n✓ Texto descifrado: %s\n", resultado)

		case 3:
			fmt.Println("\n¡Hasta luego!")
			return

		default:
			fmt.Println("\n✗ Opción inválida")
		}
	}
}

func cifrar(texto string) string {
	if len(texto) == 0 {
		return ""
	}

	// Encontrar la primera letra para calcular semilla
	var primeraLetra rune
	for _, char := range texto {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			primeraLetra = char
			break
		}
	}
	
	// Convertir a mayúscula solo para calcular semilla
	if primeraLetra >= 'a' && primeraLetra <= 'z' {
		primeraLetra = primeraLetra - 32
	}
	
	// Calcular semilla (A=1, B=2, ..., Z=26)
	semilla := int(primeraLetra - 'A' + 1)
	
	var resultado strings.Builder
	resultado.WriteString(strconv.Itoa(semilla))
	resultado.WriteString("-")

	for i, char := range texto {
		esMayuscula := char >= 'A' && char <= 'Z'
		esMinuscula := char >= 'a' && char <= 'z'
		
		if esMayuscula || esMinuscula {
			// Convertir a mayúscula temporalmente para cálculos
			charMayuscula := char
			if esMinuscula {
				charMayuscula = char - 32
			}
			
			// Posición en alfabeto (A=1, B=2, ..., Z=26)
			posicion := int(charMayuscula - 'A' + 1)
			
			// Aplicar fórmula: posición + semilla + índice
			nuevoValor := posicion + semilla + i
			
			// Aplicar módulo 26 para mantener en rango alfabético
			nuevoValor = ((nuevoValor - 1) % 26) + 1
			
			// Convertir a letra manteniendo el caso original
			if esMayuscula {
				nuevaLetra := rune('A' + nuevoValor - 1)
				resultado.WriteRune(nuevaLetra)
			} else {
				nuevaLetra := rune('a' + nuevoValor - 1)
				resultado.WriteRune(nuevaLetra)
			}
		} else {
			// Mantener caracteres no alfabéticos
			resultado.WriteRune(char)
		}
	}

	return resultado.String()
}

func descifrar(textoCifrado string) string {
	if len(textoCifrado) == 0 {
		return ""
	}

	// Separar semilla del texto cifrado
	partes := strings.SplitN(textoCifrado, "-", 2)
	if len(partes) != 2 {
		return "Error: formato inválido"
	}

	semilla, err := strconv.Atoi(partes[0])
	if err != nil {
		return "Error: semilla inválida"
	}

	textoCifrado = partes[1]
	var resultado strings.Builder

	for i, char := range textoCifrado {
		esMayuscula := char >= 'A' && char <= 'Z'
		esMinuscula := char >= 'a' && char <= 'z'
		
		if esMayuscula || esMinuscula {
			// Convertir a mayúscula temporalmente para cálculos
			charMayuscula := char
			if esMinuscula {
				charMayuscula = char - 32
			}
			
			// Posición en alfabeto de la letra cifrada
			posicionCifrada := int(charMayuscula - 'A' + 1)
			
			// Revertir la fórmula: posición original = posición cifrada - semilla - índice
			posicionOriginal := posicionCifrada - semilla - i
			
			// Ajustar si es negativo o fuera de rango
			for posicionOriginal <= 0 {
				posicionOriginal += 26
			}
			posicionOriginal = ((posicionOriginal - 1) % 26) + 1
			
			// Convertir a letra manteniendo el caso original
			if esMayuscula {
				letraOriginal := rune('A' + posicionOriginal - 1)
				resultado.WriteRune(letraOriginal)
			} else {
				letraOriginal := rune('a' + posicionOriginal - 1)
				resultado.WriteRune(letraOriginal)
			}
		} else {
			// Mantener caracteres no alfabéticos
			resultado.WriteRune(char)
		}
	}

	return resultado.String()
}