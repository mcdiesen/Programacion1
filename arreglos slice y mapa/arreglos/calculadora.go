package main

import (
	"fmt"
	"strings"
)

func sumar(expresion string) {
	valores := strings.Split(expresion, "+")
	for _, val := range valores {

	}
	fmt.Println(valores)
}

func main() {
	//llamada la funcion sumar con la expresion leida desde el teclado

	fmt.Println("==>")
	var expresion string
	fmt.Scanln(&expresion)

	sumar(expresion)

}
