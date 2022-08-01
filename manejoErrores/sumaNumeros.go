package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int
	for _, valor := range valores {
		num, err := strconv.Atoi(valor)
		//Manejo del error
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: Tiene que ingresar un numero entero")
			fmt.Println("o solo debes utilizar el operador aritmetico '+'")
		}
		result += num
	}
	return result
}
func main() {
	var expresion string
	var result int
	fmt.Println("=>")
	fmt.Scanln(&expresion)
	result = sumar(expresion)

	fmt.Println(result)
}
