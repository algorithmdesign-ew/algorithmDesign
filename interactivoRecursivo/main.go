package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== Menú de Algoritmos (Iterativo vs Recursivo) ===")
		fmt.Println("1) Búsqueda Binaria (índice en arreglo ordenado)")
		fmt.Println("2) Invertir cadena")
		fmt.Println("3) Sumar arreglo de enteros")
		fmt.Println("0) Salir")
		fmt.Print("Seleccione una opción: ")

		opStr, _ := reader.ReadString('\n')
		opStr = strings.TrimSpace(opStr)
		if opStr == "0" {
			fmt.Println("Saliendo...")
			return
		}

		method := readMethod(reader)
		switch opStr {
		case "1":
			// Binary Search
			fmt.Println("Nota: la búsqueda binaria requiere que el arreglo esté ORDENADO ascendentemente.")
			arr := readIntSlice(reader, "Ingrese enteros ORDENADOS separados por espacios o comas: ")
			target := readInt(reader, "Ingrese el valor a buscar: ")
			if method == 1 {
				start := time.Now()
				idx := binarySearchIterative(arr, target)
				dur := time.Since(start)
				if idx >= 0 {
					fmt.Printf("[Iterativo] Encontrado %d en índice %d | tiempo: %s\n\n", target, idx, dur)
				} else {
					fmt.Printf("[Iterativo] %d no encontrado | tiempo: %s\n\n", target, dur)
				}
			} else {
				start := time.Now()
				idx := binarySearchRecursive(arr, target)
				dur := time.Since(start)
				if idx >= 0 {
					fmt.Printf("[Recursivo] Encontrado %d en índice %d | tiempo: %s\n\n", target, idx, dur)
				} else {
					fmt.Printf("[Recursivo] %d no encontrado | tiempo: %s\n\n", target, dur)
				}
			}
		case "2":
			// Reverse string
			fmt.Print("Ingrese la cadena: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimRight(text, "\r\n")
			if method == 1 {
				start := time.Now()
				res := reverseStringIterative(text)
				dur := time.Since(start)
				fmt.Printf("[Iterativo] Invertida: %s | tiempo: %s\n\n", res, dur)
			} else {
				start := time.Now()
				res := reverseStringRecursive(text)
				dur := time.Since(start)
				fmt.Printf("[Recursivo] Invertida: %s | tiempo: %s\n\n", res, dur)
			}
		case "3":
			// Sum array
			arr := readIntSlice(reader, "Ingrese enteros separados por espacios o comas: ")
			if method == 1 {
				start := time.Now()
				res := sumArrayIterative(arr)
				dur := time.Since(start)
				fmt.Printf("[Iterativo] Suma = %d | tiempo: %s\n\n", res, dur)
			} else {
				start := time.Now()
				res := sumArrayRecursive(arr)
				dur := time.Since(start)
				fmt.Printf("[Recursivo] Suma = %d | tiempo: %s\n\n", res, dur)
			}
		default:
			fmt.Println("Opción inválida. Intente de nuevo.")
		}
	}
}

// ===== Utilidades de lectura =====
func readMethod(reader *bufio.Reader) int {
	for {
		fmt.Println("Seleccione el método:")
		fmt.Println("1) Iterativo")
		fmt.Println("2) Recursivo")
		fmt.Print("Opción: ")
		mStr, _ := reader.ReadString('\n')
		mStr = strings.TrimSpace(mStr)
		if mStr == "1" {
			return 1
		}
		if mStr == "2" {
			return 2
		}
		fmt.Println("Entrada inválida. Intente nuevamente.")
	}
}

func readInt(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			fmt.Println("Debe ingresar un número entero.")
			continue
		}
		v, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Formato inválido. Ejemplo válido: 42")
			continue
		}
		return v
	}
}

func readIntSlice(reader *bufio.Reader, prompt string) []int {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			fmt.Println("Debe ingresar al menos un número.")
			continue
		}
		// Reemplazar comas por espacios y dividir
		line = strings.ReplaceAll(line, ",", " ")
		parts := strings.Fields(line)
		arr := make([]int, 0, len(parts))
		ok := true
		for _, p := range parts {
			v, err := strconv.Atoi(p)
			if err != nil {
				fmt.Printf("Valor inválido: %q. Ejemplo: 1 2 3\n", p)
				ok = false
				break
			}
			arr = append(arr, v)
		}
		if ok {
			return arr
		}
	}
}

// ===== Algoritmos: Búsqueda Binaria =====
func binarySearchIterative(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarySearchRecursive(arr []int, target int) int {
	return binarySearchRecursiveRange(arr, target, 0, len(arr)-1)
}

func binarySearchRecursiveRange(arr []int, target int, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] < target {
		return binarySearchRecursiveRange(arr, target, mid+1, high)
	}
	return binarySearchRecursiveRange(arr, target, low, mid-1)
}

// ===== Algoritmos: Invertir Cadena =====
func reverseStringIterative(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseStringRecursive(s string) string {
	runes := []rune(s)
	reverseRunesRecursive(runes, 0, len(runes)-1)
	return string(runes)
}

func reverseRunesRecursive(runes []rune, left, right int) {
	if left >= right {
		return
	}
	runes[left], runes[right] = runes[right], runes[left]
	reverseRunesRecursive(runes, left+1, right-1)
}

// ===== Algoritmos: Suma de Arreglo =====
func sumArrayIterative(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func sumArrayRecursive(arr []int) int {
	return sumArrayRecursiveFrom(arr, 0)
}

func sumArrayRecursiveFrom(arr []int, idx int) int {
	if idx >= len(arr) {
		return 0
	}
	return arr[idx] + sumArrayRecursiveFrom(arr, idx+1)
}
