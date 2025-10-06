# Algoritmos Voraces en Go

Este directorio contiene dos implementaciones de algoritmos voraces desarrollados en Go para la clase de Diseño de Algoritmos.

## 📋 Contenido

- `money.go` - Algoritmo voraz para el problema de cambio de monedas
- `string.go` - Algoritmo voraz para encontrar la cadena lexicográficamente mayor

## 🪙 Algoritmo de Cambio de Monedas (`money.go`)

### Descripción
Implementa un algoritmo voraz que encuentra la combinación óptima de monedas para dar un cambio específico, minimizando el número total de monedas utilizadas.

### Funcionamiento
1. **Entrada**: Valor del cambio, cantidad de tipos de monedas, valor de cada moneda
2. **Estrategia Voraz**: Ordena las monedas de mayor a menor valor
3. **Proceso**: Toma la mayor cantidad posible de cada moneda, empezando por la de mayor valor
4. **Salida**: Conjunto C (monedas disponibles), Conjunto S (solución), y si tiene solución

### Ejemplo de Uso
```
Valor del cambio: 67
Cantidad de tipos de monedas: 4
Valores: 50, 20, 10, 5

Resultado:
Conjunto C: {50, 20, 10, 5}
Conjunto S: {1×50, 1×10, 1×5, 1×2} // Error aquí, debería ser {1×50, 0×20, 1×10, 1×5}
```

### Aspectos Técnicos de Go

#### `sort.Sort(sort.Reverse(sort.IntSlice(monedas)))`
- **`sort.IntSlice`**: Convierte el slice de enteros en un tipo que implementa la interfaz `sort.Interface`
- **`sort.Reverse`**: Invierte el orden de comparación para ordenar de mayor a menor
- **Complejidad**: O(n log n)

#### `map[int]int` para conjuntoS
```go
conjuntoS := make(map[int]int) // moneda -> cantidad usada
```
- **Map**: Estructura de datos hash que permite acceso O(1) promedio
- **Clave**: Valor de la moneda
- **Valor**: Cantidad de esa moneda utilizada

#### Manejo de Errores
```go
valorCambio, err := strconv.Atoi(scanner.Text())
if err != nil || valorCambio < 0 {
    fmt.Println("Error: Valor inválido")
    return
}
```
- **`strconv.Atoi`**: Convierte string a entero, retorna valor y error
- **Validación**: Verifica tanto error de conversión como valores negativos

## 📝 Algoritmo Lexicográfico (`string.go`)

### Descripción
Encuentra la siguiente permutación lexicográficamente mayor de una cadena de caracteres usando una estrategia voraz.

### Funcionamiento
1. **Entrada**: Cadena de caracteres
2. **Estrategia Voraz**: Busca el punto de intercambio más a la derecha posible
3. **Proceso**: 
   - Encuentra el carácter más a la derecha que es menor que su sucesor
   - Intercambia con el menor carácter mayor a su derecha
   - Ordena la parte derecha en orden descendente
4. **Salida**: Cadena lexicográficamente mayor

### Ejemplo de Uso
```
Entrada: "abc"
Salida: "acb"

Entrada: "321"
Salida: "321" (ya es la mayor posible)
```

### Aspectos Técnicos de Go

#### Conversión a Runes
```go
runes := []rune(s)
```
- **`[]rune`**: Maneja correctamente caracteres Unicode
- **Diferencia con `[]byte`**: Rune representa un punto de código Unicode completo
- **Importancia**: Permite trabajar con caracteres especiales y emojis correctamente

#### Búsqueda del Punto de Intercambio
```go
i := n - 2
for i >= 0 && runes[i] >= runes[i+1] {
    i--
}
```
- **Algoritmo**: Recorre de derecha a izquierda buscando el primer descenso
- **Complejidad**: O(n) en el peor caso
- **Estrategia Voraz**: Encuentra el punto más a la derecha para maximizar el resultado

#### Ordenamiento con Slice Personalizado
```go
sort.Slice(derecha, func(a, b int) bool {
    return derecha[a] > derecha[b]
})
```
- **`sort.Slice`**: Ordena usando una función de comparación personalizada
- **Función anónima**: Define el criterio de ordenamiento (descendente)
- **Ventaja**: No requiere implementar la interfaz `sort.Interface`

#### Intercambio de Valores
```go
runes[i], runes[j] = runes[j], runes[i]
```
- **Intercambio múltiple**: Característica nativa de Go para intercambiar valores
- **Sin variables temporales**: Más limpio que el patrón tradicional
- **Eficiencia**: Compilador optimiza esta operación

## 🚀 Ejecución

```bash
# Ejecutar algoritmo de monedas
go run money.go

# Ejecutar algoritmo lexicográfico  
go run string.go
```

## 📊 Complejidad Computacional

### Algoritmo de Monedas
- **Tiempo**: O(n log n) por el ordenamiento + O(n) por el procesamiento = O(n log n)
- **Espacio**: O(n) para almacenar las monedas y la solución

### Algoritmo Lexicográfico
- **Tiempo**: O(n) para encontrar el punto + O(k log k) para ordenar la parte derecha = O(n log n)
- **Espacio**: O(n) para el slice de runes

## 🎯 Características de la Estrategia Voraz

1. **Elección Voraz**: En cada paso se toma la decisión localmente óptima
2. **Sin Retroceso**: No se reconsidera ninguna decisión tomada
3. **Eficiencia**: Generalmente más rápidos que algoritmos de programación dinámica
4. **Limitaciones**: No siempre garantiza la solución globalmente óptima

## 🔧 Dependencias

- Go 1.16 o superior
- Paquetes estándar: `bufio`, `fmt`, `os`, `sort`, `strconv`