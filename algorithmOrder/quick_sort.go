package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// leerCSV: Lee un archivo CSV y retorna un slice de números enteros
func leerCSV(nombreArchivo string) ([]int, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo: %v", err)
	}
	defer archivo.Close()

	var numeros []int
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}
		valores := strings.Split(linea, ",")
		for _, valor := range valores {
			valor = strings.TrimSpace(valor)
			if valor == "" {
				continue
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

// quickSort: Se implementa el algoritmo de ordenamiento rápido (in-place)
// Estrategia: particionar el array alrededor de un pivote y ordenar recursivamente
func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSortRange(nums, 0, len(nums)-1)
}

func quickSortRange(nums []int, low, high int) {
	for low < high {
		p := partition(nums, low, high)
		// Optimización tail-recursion: ordenar primero la sublista más pequeña
		if p-low < high-p {
			quickSortRange(nums, low, p-1)
			low = p + 1
		} else {
			quickSortRange(nums, p+1, high)
			high = p - 1
		}
	}
}

// partition particiona el slice usando el último elemento como pivote
func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return i
}

func guardarResultado(nombreArchivo string, numeros []int) error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("error creando archivo: %v", err)
	}
	defer archivo.Close()

	for i, numero := range numeros {
		_, err = fmt.Fprintf(archivo, "%d", numero)
		if err != nil {
			return err
		}
		if i < len(numeros)-1 {
			_, err = fmt.Fprintf(archivo, ",")
			if err != nil {
				return err
			}
		}
	}
	_, err = fmt.Fprintf(archivo, "\n")
	return err
}

func mostrarEstadisticas(numeros []int, tiempoEjecucion time.Duration) {
	fmt.Println("\n📊 Estadísticas del Ordenamiento (Quick Sort):")
	fmt.Printf("   Cantidad de números: %d\n", len(numeros))
	fmt.Printf("   Tiempo de ejecución: %v\n", tiempoEjecucion)
	fmt.Printf("   Primeros 5 ordenados: %v\n", numeros[:5])
	fmt.Printf("   Últimos 5 ordenados: %v\n", numeros[len(numeros)-5:])

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

func main() {
	fmt.Println("🔢 Algoritmo de Ordenamiento - Quick Sort")
	fmt.Println("==============================================")

	archivoEntrada := "numeros_10000.csv"
	if len(os.Args) > 1 && strings.TrimSpace(os.Args[1]) != "" {
		archivoEntrada = os.Args[1]
	}
	fmt.Printf("\n📁 Leyendo archivo: %s\n", archivoEntrada)

	numeros, err := leerCSV(archivoEntrada)
	if err != nil {
		fmt.Printf("❌ Error leyendo archivo: %v\n", err)
		return
	}
	fmt.Printf("   Números leídos: %d\n", len(numeros))
	fmt.Printf("   Primeros 5 originales: %v\n", numeros[:5])

	numerosACopiar := make([]int, len(numeros))
	copy(numerosACopiar, numeros)

	fmt.Println("\n🔄 Iniciando ordenamiento con Quick Sort...")
	inicio := time.Now()
	quickSort(numerosACopiar)
	tiempoEjecucion := time.Since(inicio)

	mostrarEstadisticas(numerosACopiar, tiempoEjecucion)

	archivoSalida := "numeros_10000_ordenados_quick.csv"
	if len(os.Args) > 1 && strings.TrimSpace(os.Args[1]) != "" {
		base := strings.TrimSuffix(os.Args[1], ".csv")
		archivoSalida = base + "_ordenados_quick.csv"
	}
	if err := guardarResultado(archivoSalida, numerosACopiar); err != nil {
		fmt.Printf("❌ Error guardando resultado: %v\n", err)
	} else {
		fmt.Printf("\n💾 Resultado guardado en: %s\n", archivoSalida)
	}

	fmt.Println("\n✅ Proceso completado exitosamente!")
}
