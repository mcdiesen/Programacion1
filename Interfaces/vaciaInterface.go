package main

import "fmt"

//interface vacia no tiene métodos y puede ser implementado por cualquier tipo

func main() {
	//aux Interface implementada por nil, int, string
	var aux interface{}
	fmt.Println(aux)
	aux = 10
	fmt.Println(aux)

	aux = "Hello"

	fmt.Println(aux)
}

//La salida del programa sera:
//nil
//10
//Hello
