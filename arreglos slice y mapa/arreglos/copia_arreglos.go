package main

import "fmt"

func main() {
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
	//copia el arreglo planets
	planetsMarkII := planets
	//Agrega el string whoops en la posición 2 del arreglo
	planets[2] = "whoops"
	//imprime el arreglo planets y planetsMarkII
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
