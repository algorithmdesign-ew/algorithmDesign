# interactivoRecursivo

Este módulo implementa 3 problemas, cada uno resuelto de dos formas: iterativa y recursiva. El objetivo es comparar ambos enfoques para el mismo problema.

Algoritmos incluidos (Iterativo y Recursivo):
- Búsqueda Binaria (índice en arreglo ordenado)
  - Iterativo: `binarySearchIterative(arr []int, target int) int`
  - Recursivo: `binarySearchRecursive(arr []int, target int) int`
- Invertir cadena (soporta Unicode con `[]rune`)
  - Iterativo: `reverseStringIterative(s string) string`
  - Recursivo: `reverseStringRecursive(s string) string`
- Sumar arreglo de enteros
  - Iterativo: `sumArrayIterative(arr []int) int`
  - Recursivo: `sumArrayRecursive(arr []int) int`

Sigue el menú para elegir el problema y el método (iterativo o recursivo).

## Iterativo vs Recursivo (resumen práctico)
- Iterativo: usa bucles (`for`), evita costo de llamadas y consumo de pila. Suele ser más rápido y seguro para entradas muy grandes en Go.
- Recursivo: divide el problema en subproblemas con casos base. Es más directo en problemas naturalmente recursivos (árboles, backtracking), pero consume pila por cada llamada.

Regla rápida en Go:
- Preferir iterativo para procesamiento lineal de arreglos, transformaciones simples y entradas grandes.
- Considerar recursivo para estructuras jerárquicas o divide & conquer, validando tamaños para no desbordar la pila.

## Detalle por algoritmo

### 1) Búsqueda Binaria
- Problema: Dado un arreglo ORDENADO ascendentemente y un objetivo `target`, devolver el índice o `-1` si no existe.
- Complejidad: O(log n) tiempo, O(1) memoria iterativa; O(log n) memoria recursiva (pila).
- Iterativo: mantiene `low` y `high`, calcula `mid`, compara y reduce el rango.
- Recursivo: mismo principio, llamando sobre el subrango `low..mid-1` o `mid+1..high`.

### 2) Invertir cadena
- Problema: Devolver la cadena invertida.
- Complejidad: O(n) tiempo. Memoria: O(1) iterativa; O(n) recursiva (pila).
- Detalle: Se usa `[]rune` para soportar Unicode correctamente.

### 3) Sumar arreglo
- Problema: Sumar todos los elementos de un `[]int`.
- Complejidad: O(n) tiempo. Memoria: O(1) iterativa; O(n) recursiva.

## Entradas
- Búsqueda Binaria: el arreglo debe estar ordenado ascendentemente.
- Invertir cadena: cualquier cadena (Unicode soportado).
- Sumar arreglo: lista de enteros separada por espacios o comas.

## Notas
- En Go no hay tail-call optimization. Para entradas grandes, la versión iterativa evita desbordar la pila.
- El código separa lectura de entradas y lógica de cálculo para mayor claridad y prueba.
