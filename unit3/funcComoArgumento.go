//Funciones como argumentos
//Indica el operador y los operandos
package main

import "fmt"

//Doit funcion que recibe una función como parámetro.
func doit(operator func(int, int) int, a int, b int) int {
	return operator(a, b)
}
func sum(a int, b int) int {
	return a + b
}
func multiply(a int, b int) int {
	return a * b
}
func main() {
	c := doit(sum, 2, 3)
	fmt.Println("2+3=", c)
	d := doit(multiply, 2, 3)
	fmt.Println("2*3=", d)
}

/*
run the program: funcComoArgumento.go
Output:
	2+3=5
	2*3=6
*/
