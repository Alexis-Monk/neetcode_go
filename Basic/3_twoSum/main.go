package main

import "fmt"

func twoSum(nums []int, target int) []int { // se crea la func la cual se le van a designar un array de ints y un target int, devolverá un array de ints
	prevMap := make(map[int]int) // se crea un map con clave int y valor int

	for i, n := range nums { //se itera en nums tanto indices como valores
		diff := target - n                // se genera la diff del target y numero iterado para tener el valor de la diferencia
		if val, ok := prevMap[diff]; ok { // se verifica si el valor se encuentre en el prevmap, si existe
			return []int{val, i} // regresa el slice indice que se tenia en el prevmap, el indice actual
		}
		prevMap[n] = i // sino agrega el indice al prevmap
	}
	return nil // sino regresa nil
}

func twoSum2(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if val, ok := prevMap[diff]; ok {
			return []int{val, i}
		}
		prevMap[n] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum2(nums, target)
	fmt.Println(result)
}
