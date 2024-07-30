package main

import "fmt"

func containsDuplicate(nums []int) bool { // se crea la función a la cual se le van a asignar un array de numeros enteros y regresará un bool
	hashset := make(map[int]bool) // se va a crear un mapa la cual la clave será un int y valor bool

	for _, n := range nums { // se hará el loop de la lista
		if hashset[n] { // si le numero se encuentra en el mapa
			return true // retorna true
		}
		hashset[n] = true // si no está agrega el valor al hashset
	}
	return false // si acaba el loop y no encontró regresa false
}

// 1. Create a map to simulate a set where we store unique numbers.
// 2. Iterate over the slice of integers.
// 3. Check if the number is already present in the set.
// 4. If a duplicate is found, return true.
// 5. Add the number to the set.
// 6. If the loop completes without finding duplicates, return false.

func containsDuplicate2(nums []int) bool {
	hashset := make(map[int]bool)

	for _, n := range nums {
		if hashset[n] {
			return true
		}
		hashset[n] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate2(nums))
}
