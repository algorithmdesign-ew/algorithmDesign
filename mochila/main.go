package main

import (
	"fmt"
)

type Articulo struct {
	nombre string
	peso   int
	valor  int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mochila(articulos []Articulo, capacidad int) (int, []string) {
	n := len(articulos)
	
	// Crear tabla de programación dinámica
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacidad+1)
	}
	
	// Llenar la tabla
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacidad; w++ {
			// Si el artículo no cabe, no lo incluimos
			if articulos[i-1].peso > w {
				dp[i][w] = dp[i-1][w]
			} else {
				// Tomar el máximo entre incluir o no incluir el artículo
				dp[i][w] = max(dp[i-1][w], 
					dp[i-1][w-articulos[i-1].peso]+articulos[i-1].valor)
			}
		}
	}
	
	// Reconstruir la solución (qué artículos se incluyeron)
	articulosSeleccionados := []string{}
	w := capacidad
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			articulosSeleccionados = append(articulosSeleccionados, articulos[i-1].nombre)
			w -= articulos[i-1].peso
		}
	}
	
	return dp[n][capacidad], articulosSeleccionados
}

func main() {
	var numArticulos, capacidadMochila int
	
	fmt.Println("=== PROBLEMA DE LA MOCHILA ===")
	fmt.Println()
	
	// Preguntar capacidad de la mochila
	fmt.Print("Ingrese la capacidad de la mochila: ")
	fmt.Scan(&capacidadMochila)
	
	// Preguntar número de artículos
	fmt.Print("¿Cuántos artículos hay?: ")
	fmt.Scan(&numArticulos)
	fmt.Println()
	
	articulos := make([]Articulo, numArticulos)
	
	// Leer información de cada artículo
	for i := 0; i < numArticulos; i++ {
		var peso, valor int
		nombre := fmt.Sprintf("Artículo %d", i+1)
		
		fmt.Printf("Artículo %d:\n", i+1)
		fmt.Print("  Peso (Xi): ")
		fmt.Scan(&peso)
		fmt.Print("  Valor (Vi): ")
		fmt.Scan(&valor)
		
		articulos[i] = Articulo{
			nombre: nombre,
			peso:   peso,
			valor:  valor,
		}
		fmt.Println()
	}
	
	// Resolver el problema
	valorMaximo, articulosSeleccionados := mochila(articulos, capacidadMochila)
	
	// Mostrar resultado
	fmt.Println("=== RESULTADO ===")
	fmt.Printf("El viajero con una mochila de peso %d debe llevar los artículos: ", capacidadMochila)
	
	// Invertir el orden para mostrar correctamente
	for i := len(articulosSeleccionados) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%s", articulosSeleccionados[i])
		} else {
			fmt.Printf("%s, ", articulosSeleccionados[i])
		}
	}
	
	fmt.Printf(" con una ganancia de %d\n", valorMaximo)
}