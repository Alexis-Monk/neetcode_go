package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) { // Se declara la función donde son 2 arrays llamados nums y las variables n, m, indican la cantidad de números que tienen los arrays
	for n != 0 { // inicia el bucle for siempre y cuando el valor de n sea distinto de 0
		if m != 0 && nums1[m-1] > nums2[n-1] { // si el valor de n es distinto de 0 y el último valor de nums1 es mayor al de nums2
			nums1[m+n-1] = nums1[m-1] // al ultimo valor del array completo se le asigna el último valor del array nums1
			m--                       // al valor de m se le resta uno
		} else { // si o es el caso del if
			nums1[m+n-1] = nums2[n-1] // al último valor del array completo de nums1 se le asigna el último valor del array de nums2
			n--                       // se le resta 1 al valor de n
		}
	}
}

/*




Espacio para no visualizar la función de arriba




*/

func merge2(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge2(nums1, m, nums2, n)
	fmt.Println(nums1)
}
