package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// GeneradorRandom estructura simple para generar números únicos
type GeneradorRandom struct {
	generador     *rand.Rand
	arrayCompleto []int
	indiceActual  int
}

// crea un nuevo generador
func NewGeneradorRandom() *GeneradorRandom {
	// Crear semilla única basada en tiempo
	source := rand.NewSource(time.Now().UnixNano())

	return &GeneradorRandom{
		generador:     rand.New(source),
		arrayCompleto: make([]int, 10000),
		indiceActual:  0,
	}
}

// inicializar. Llena el array con números del 1 al 10000 y los mezcla
func (g *GeneradorRandom) inicializar() {
	// Paso 1: Llenar array con números ordenados (1, 2, 3, ..., 10000)
	for i := 0; i < 10000; i++ {
		g.arrayCompleto[i] = i + 1
	}

	// Paso 2: Mezclar usando Fisher-Yates
	for i := 9999; i > 0; i-- {
		j := g.generador.Intn(i + 1)
		g.arrayCompleto[i], g.arrayCompleto[j] = g.arrayCompleto[j], g.arrayCompleto[i]
	}

	g.indiceActual = 0
}

// GenerarNumeros genera 'cantidad' números únicos del 1 al 10000
func (g *GeneradorRandom) GenerarNumeros(cantidad int) []int {
	// Validar que no exceda el límite
	if cantidad > 10000 {
		panic("No se pueden generar más de 10000 números únicos")
	}

	// Si es la primera vez, inicializar
	if g.indiceActual == 0 {
		g.inicializar()
	}

	// Si no hay suficientes números, reinicializar
	if g.indiceActual+cantidad > 10000 {
		g.inicializar()
	}

	// Extraer los números solicitados
	resultado := make([]int, cantidad)
	copy(resultado, g.arrayCompleto[g.indiceActual:g.indiceActual+cantidad])
	g.indiceActual += cantidad

	return resultado
}

// guarda los números en un archivo .csv
func guardarEnArchivo(nombreArchivo string, numeros []int, descripcion string) error {
	// Crear archivo (sobrescribir si existe)
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return fmt.Errorf("error creando archivo: %v", err)
	}
	defer archivo.Close()

	// Escribir números en formato CSV
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
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Crear generador
	generador := NewGeneradorRandom()

	fmt.Println("🎲 Generador de Números Random (1-10000)")
	fmt.Println("==========================================")

	// Generar 10000 números (ideal para algoritmos de ordenamiento)
	fmt.Println("\n📊 Generando 10000 números únicos:")
	numeros := generador.GenerarNumeros(10000)
	fmt.Printf("   Cantidad generada: %d\n", len(numeros))
	fmt.Printf("   Primeros 5: %v\n", numeros[:5])
	fmt.Printf("   Últimos 5: %v\n", numeros[995:])

	// Guardar en archivo fmt CSV
	err := guardarEnArchivo("numeros_10000.csv", numeros, "10000 números únicos para algoritmos de ordenamiento")
	if err != nil {
		fmt.Printf("❌ Error guardando archivo: %v\n", err)
	} else {
		fmt.Printf("💾 Guardado en: numeros_10000.csv\n")
	}

	// Mostrar estadísticas en consola
	min, max := numeros[0], numeros[0]
	suma := 0
	for _, num := range numeros {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		suma += num
	}
	promedio := float64(suma) / float64(len(numeros))

	fmt.Println("\n📈 Estadísticas de los números generados:")
	fmt.Printf("   Mínimo: %d\n", min)
	fmt.Printf("   Máximo: %d\n", max)
	fmt.Printf("   Promedio: %.2f\n", promedio)
	fmt.Printf("   Rango: %d\n", max-min)

	fmt.Println("\n✅ Generador listo para algoritmos de ordenamiento!")
	fmt.Println("📁 Archivo creado: numeros_1000.csv")
}
