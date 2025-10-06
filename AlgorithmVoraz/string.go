package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	
	fmt.Println("\n--- RESULTADO ---")
	fmt.Println("Cadena original:", cadena)
	
	// Algoritmo voraz: encontrar la cadena lexicográficamente mayor
	resultado := encontrarCadenaLexicograficaMayor(cadena)
	
	fmt.Println("Cadena lexicográficamente MAYOR:", resultado)
}

func encontrarCadenaLexicograficaMayor(s string) string {
	runes := []rune(s)
	n := len(runes)
	
	if n <= 1 {
		return s
	}
	
	// Estrategia voraz: buscar de derecha a izquierda el primer carácter
	// que sea menor que algún carácter a su derecha
	i := n - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}
	
	// Si no encontramos tal carácter, la cadena ya está en orden descendente
	if i < 0 {
		return s // Ya es la mayor posible
	}
	
	// Encontrar el carácter más pequeño a la derecha de i que sea mayor que runes[i]
	j := n - 1
	for j > i && runes[j] <= runes[i] {
		j--
	}
	
	// Intercambiar
	runes[i], runes[j] = runes[j], runes[i]
	
	// Ordenar la parte derecha en orden descendente para maximizar
	derecha := runes[i+1:]
	sort.Slice(derecha, func(a, b int) bool {
		return derecha[a] > derecha[b]
	})
	
	return string(runes)
}