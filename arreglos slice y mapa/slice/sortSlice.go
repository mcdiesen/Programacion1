package main

import (
	"fmt"
	"sort"
)

func main() {
	//crear un slice con los nombre de los planetas
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	//Ordena los planetas en orden alfabético
	sort.StringSlice(planets).Sort()

	//El siguiente código es equivalente: sort.String(planets)

	//imprime los elementos del slice
	fmt.Println(planets)
}

//La salida: [Earth Jupiter Mars Mercury Neptune Saturn Uranus Venus]
