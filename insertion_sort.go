package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// leerCSV lee un archivo CSV y retorna un slice de números enteros
func leerCSV(nombreArchivo string) ([]int, error) {
	// Abrir archivo
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo: %v", err)
	}
	defer archivo.Close()

	var numeros []int
	scanner := bufio.NewScanner(archivo)

	// Leer línea por línea
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue // Saltar líneas vacías
		}

		// Dividir por comas
		valores := strings.Split(linea, ",")

		// Convertir cada valor a entero
		for _, valor := range valores {
			valor = strings.TrimSpace(valor)
			if valor == "" {
				continue // Saltar valores vacíos
			}

			numero, err := strconv.Atoi(valor)
			if err != nil {
				return nil, fmt.Errorf("error convirtiendo '%s' a número: %v", valor, err)
			}

			numeros = append(numeros, numero)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error leyendo archivo: %v", err)
	}

	return numeros, nil
}

// insertionSort: Se implementa el algoritmo de ordenamiento por inserción
// Explicación del algoritmo:
// 1. Comenzamos con el segundo elemento (índice 1)
// 2. Para cada elemento, lo comparamos con los anteriores
// 3. Si es menor, lo movemos hacia la izquierda
// 4. Repetimos hasta que esté en su posición correcta
func insertionSort(numeros []int) {
	n := len(numeros)

	// Iterar desde el segundo elemento hasta el final
	for i := 1; i < n; i++ {
		// Guardar el elemento actual
		elementoActual := numeros[i]

		// j es la posición donde insertaremos el elemento
		j := i - 1

		// Mover elementos mayores hacia la derecha
		// hasta encontrar la posición correcta
		for j >= 0 && numeros[j] > elementoActual {
			numeros[j+1] = numeros[j]
			j--
		}

		// Insertar el elemento en su posición correcta
		numeros[j+1] = elementoActual
	}
}

// mostrarEstadisticas muestra información sobre el array ordenado
func mostrarEstadisticas(numeros []int, tiempoEjecucion time.Duration) {
	fmt.Println("\n📊 Estadísticas del Ordenamiento:")
	fmt.Printf("   Cantidad de números: %d\n", len(numeros))
	fmt.Printf("   Tiempo de ejecución: %v\n", tiempoEjecucion)
	fmt.Printf("   Primeros 5 ordenados: %v\n", numeros[:5])
	fmt.Printf("   Últimos 5 ordenados: %v\n", numeros[len(numeros)-5:])

	// Verificar que esté ordenado correctamente
	ordenado := true
	for i := 1; i < len(numeros); i++ {
		if numeros[i] < numeros[i-1] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println("   ✅ Verificación: Array correctamente ordenado")
	} else {
		fmt.Println("   ❌ Verificación: Array NO está ordenado")
	}
}

// guardarResultado guarda el array ordenado en un nuevo archivo CSV
func guardarResultado(nombreArchivo string, numeros []int) error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("error creando archivo: %v", err)
	}
	defer archivo.Close()

	// Escribir números ordenados en formato CSV
	for i, numero := range numeros {
		_, err = fmt.Fprintf(archivo, "%d", numero)
		if err != nil {
			return err
		}

		// Agregar coma, excepto para el último número
		if i < len(numeros)-1 {
			_, err = fmt.Fprintf(archivo, ",")
			if err != nil {
				return err
			}
		}
	}

	// Agregar nueva línea al final
	_, err = fmt.Fprintf(archivo, "\n")
	return err
}

func main() {
	fmt.Println("🔢 Algoritmo de Ordenamiento - Insertion Sort")
	fmt.Println("==============================================")

	// Nombre del archivo CSV a procesar
	archivoEntrada := "numeros_10000.csv"

	fmt.Printf("\n📁 Leyendo archivo: %s\n", archivoEntrada)

	// Leer números del CSV
	numeros, err := leerCSV(archivoEntrada)
	if err != nil {
		fmt.Printf("❌ Error leyendo archivo: %v\n", err)
		return
	}

	fmt.Printf("   Números leídos: %d\n", len(numeros))
	fmt.Printf("   Primeros 5 originales: %v\n", numeros[:5])

	// Crear una copia para no modificar el original
	numerosACopiar := make([]int, len(numeros))
	copy(numerosACopiar, numeros)

	fmt.Println("\n🔄 Iniciando ordenamiento con Insertion Sort...")

	// Medir tiempo de ejecución
	inicio := time.Now()
	insertionSort(numerosACopiar)
	tiempoEjecucion := time.Since(inicio)

	// Mostrar estadísticas
	mostrarEstadisticas(numerosACopiar, tiempoEjecucion)

	// Guardar resultado ordenado
	archivoSalida := "numeros_1000_ordenados.csv"
	err = guardarResultado(archivoSalida, numerosACopiar)
	if err != nil {
		fmt.Printf("❌ Error guardando resultado: %v\n", err)
	} else {
		fmt.Printf("\n💾 Resultado guardado en: %s\n", archivoSalida)
	}

	fmt.Println("\n✅ Proceso completado exitosamente!")
}
