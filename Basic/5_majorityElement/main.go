package main

import "fmt"

func majorityElement(nums []int) int { // Función que toma el parámetro array nums y devuelve un int
	m := map[int]int{} // Se crea el mapa para contar la frecuencia de cada número en nums

	for _, n := range nums { // Se va a iterar sobre nums
		m[n]++                  // donde para cada elemento n, se va a sumar 1
		if m[n] > len(nums)/2 { // cuenta si el contador del número actual n en el mapa es mayor que la mitad de la longitud de nms
			return n // regresa n
		}
	}

	return -1 // Si ningun elemento se encontró con mayor frecuencia mayor a la mitad del arreglo, regresa -1 que indica que no hay un elemento mayoritario
}

func majorityElement2(nums []int) int {
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
		if m[n] > len(nums)/2 {
			return n
		}
	}

	return -1
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	result := majorityElement2(nums)
	fmt.Println(result)
}
