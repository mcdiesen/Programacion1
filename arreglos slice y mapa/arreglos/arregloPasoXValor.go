package main

import "fmt"

//La función terraform no afecta al arreglo original
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}
func main() {
	//Arreglo planets en la funcion main
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	//paso por valor del arreglo planets a la función terraform
	terraform(planets)
	fmt.Println(planets)
}
