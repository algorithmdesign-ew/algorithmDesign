package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║  ALGORITMO VORAZ: CAMBIO DE MONEDAS       ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println()
	
	// Leer valor del cambio
	fmt.Print("Ingrese el valor del cambio: ")
	scanner.Scan()
	valorCambio, err := strconv.Atoi(scanner.Text())
	if err != nil || valorCambio < 0 {
		fmt.Println("Error: Valor inválido")
		return
	}
	
	// Leer cantidad de monedas
	fmt.Print("Ingrese la cantidad de tipos de monedas: ")
	scanner.Scan()
	cantidadMonedas, err := strconv.Atoi(scanner.Text())
	if err != nil || cantidadMonedas <= 0 {
		fmt.Println("Error: Cantidad inválida")
		return
	}
	
	// Leer valores de cada moneda
	monedas := make([]int, cantidadMonedas)
	for i := 0; i < cantidadMonedas; i++ {
		fmt.Printf("Ingrese el valor de la moneda %d: ", i+1)
		scanner.Scan()
		monedas[i], err = strconv.Atoi(scanner.Text())
		if err != nil || monedas[i] <= 0 {
			fmt.Println("Error: Valor de moneda inválido")
			return
		}
	}
	
	// Ordenar monedas de mayor a menor (estrategia voraz)
	sort.Sort(sort.Reverse(sort.IntSlice(monedas)))
	
	// Conjunto C (monedas disponibles ordenadas)
	fmt.Println("\n--- RESULTADO ---")
	fmt.Print("Conjunto C (monedas disponibles): {")
	for i, m := range monedas {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(m)
	}
	fmt.Println("}")
	
	// Algoritmo voraz para encontrar el cambio
	conjuntoS := make(map[int]int) // moneda -> cantidad usada
	cambioRestante := valorCambio
	
	for _, moneda := range monedas {
		if cambioRestante >= moneda {
			cantidad := cambioRestante / moneda
			conjuntoS[moneda] = cantidad
			cambioRestante -= cantidad * moneda
		}
	}
	
	// Conjunto S (solución)
	fmt.Print("Conjunto S (solución): {")
	primero := true
	for _, moneda := range monedas {
		if cant, existe := conjuntoS[moneda]; existe {
			if !primero {
				fmt.Print(", ")
			}
			fmt.Printf("%d×%d", cant, moneda)
			primero = false
		}
	}
	fmt.Println("}")
	
	// Verificar si hay solución
	if cambioRestante == 0 {
		fmt.Println("\n✓ TIENE SOLUCIÓN")
		fmt.Printf("Total de monedas usadas: %d\n", contarMonedas(conjuntoS))
	} else {
		fmt.Println("\n✗ NO TIENE SOLUCIÓN")
		fmt.Printf("Faltante: %d\n", cambioRestante)
	}
}

func contarMonedas(conjunto map[int]int) int {
	total := 0
	for _, cant := range conjunto {
		total += cant
	}
	return total
}