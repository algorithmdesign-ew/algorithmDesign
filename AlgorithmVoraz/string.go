package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║  ALGORITMO VORAZ: LEXICOGRÁFICAMENTE MAYOR ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println()

	// Leer cadena de entrada
	fmt.Print("Ingrese la cadena de caracteres: ")
	scanner.Scan()
	cadena := scanner.Text()

	if len(cadena) == 0 {
		fmt.Println("Error: Cadena vacía")
		return
	}

	fmt.Println("\n--- PROCESO DEL ALGORITMO VORAZ ---")
	fmt.Println("Cadena original:", cadena)
	fmt.Println()

	// Algoritmo voraz: encontrar la cadena lexicográficamente mayor
	resultado := encontrarCadenaLexicograficaMayor(cadena)

	fmt.Println("\n--- RESULTADO FINAL ---")
	fmt.Println("Cadena lexicográficamente MAYOR:", resultado)
}

func encontrarCadenaLexicograficaMayor(s string) string {
	runes := []rune(s)
	n := len(runes)

	if n <= 1 {
		return s
	}

	fmt.Println("ESTRATEGIA VORAZ: Cada carácter se mueve al INICIO si es mayor")
	fmt.Println("que cualquier carácter a su izquierda\n")

	// Procesar cada carácter de izquierda a derecha
	for i := 1; i < n; i++ {
		caracterActual := runes[i]
		fmt.Printf("Paso %d - Analizando carácter '%c' en posición %d\n", i, caracterActual, i)
		
		// Verificar si es mayor que ALGÚN carácter a su izquierda
		esMayor := false
		for j := 0; j < i; j++ {
			if caracterActual > runes[j] {
				esMayor = true
				break
			}
		}
		
		if esMayor {
			// Mover el carácter al INICIO (posición 0)
			fmt.Printf("  → '%c' es MAYOR que algún carácter a su izquierda\n", caracterActual)
			fmt.Printf("  → Decisión voraz: Mover '%c' al INICIO\n", caracterActual)
			
			// Guardar el carácter
			temp := runes[i]
			// Desplazar todos los caracteres anteriores una posición a la derecha
			for k := i; k > 0; k-- {
				runes[k] = runes[k-1]
			}
			// Colocar el carácter al inicio
			runes[0] = temp
			
			fmt.Printf("  → Cadena actual: %s\n\n", string(runes))
		} else {
			fmt.Printf("  → '%c' NO es mayor que ningún carácter a su izquierda\n", caracterActual)
			fmt.Printf("  → Se queda en su posición\n")
			fmt.Printf("  → Cadena actual: %s\n\n", string(runes))
		}
	}

	return string(runes)
}