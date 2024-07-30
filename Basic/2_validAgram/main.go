package main

import "fmt"

func isAnagram(s1, t string) bool { // se crea la func que se le asignarán dos strings y devolverá un bool
	if len(s1) != len(t) { // se las longitudes de ambos strings son diferentes
		return false // returna false
	}

	countS, countT := make(map[rune]int), make(map[rune]int) // se generan dos mapas para ambos strings que van a tener una key rune y valor int

	for _, char := range s1 { // se itera a través de s1 tomando en cuenta el valor
		countS[char]++ // cada vez que salga el character, el mapa aumentará uno
	}

	for _, char := range t { // se itera en t
		countT[char]++ // cada vez que salga el valor se agrega uno en el map
	}

	for char, count := range countS { // se iteran los caracteres en el mapa
		if countT[char] != count { // si el conteo en el mapa T es diferente al de S
			return false // regresa false
		}
	}

	return true // sino regresa true
}

/*





Espacio para no visualizar la función de arriba





*/

func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)

	for _, char := range s {
		countS[char]++
	}

	for _, char := range t {
		countT[char]++
	}

	for char, count := range countS {
		if countT[char] != count {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram2(s, t))
}
