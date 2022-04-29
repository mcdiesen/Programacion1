//declaracion de variables en Go
package main

import "fmt"

func Variable() {
	//Declaracion de variables con tipos de datos básicos
	var (
		nombre  string
		edad    int
		estudia bool
		salario float64
	)
	//Declaracion de variable en formato corto
	nivelAcad := "Ingeniería"
	//asignación de valores a las variables
	nombre = "Charles McLean"
	edad = 39
	estudia = true
	salario = 1000

	//imprimir los valores de cada variable con la funcion Println
	fmt.Println("Nombre: ", nombre)
	fmt.Println("edad: ", edad)
	fmt.Println("estudia: ", estudia)
	fmt.Println("Salario: ", salario)
	fmt.Println("Nivel Académico: ", nivelAcad)
}
