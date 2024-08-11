package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int { // declara s como un string y devuelve un entero
	words := strings.Fields(s) // Separa la cadena en un slice de palabras usando los espacios como delimitadores
	if len(words) == 0 {       // Si no se tiene ninguna palabra
		return 0 // devuelve 0
	}
	return len(words[len(words)-1]) // Retorna la longitud de la última palabra
}

func lengthOfLastWord2(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	return len(words[len(words)-1])
}

func main() {
	s := "fly me to "
	result := lengthOfLastWord2(s)
	fmt.Println(result)
}
