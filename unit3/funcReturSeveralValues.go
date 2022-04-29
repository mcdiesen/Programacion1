//Función que devuelve múltiples valores
package main

import "fmt"

//Ops function return two values. Adds two numbers, minus two numbers
func ops(a int, b int) (int, int) {
	return a + b, a - b
}
func main() {
	sum, subs := ops(2, 2)
	fmt.Println("2+2=", sum, "2-2=", subs)
	b, _ := ops(10, 2)
	fmt.Println("10+2=", b)
}

/*
run the program: go run funcReturSeveralValues.go
output:
2+2=4	2-2=0
10+2=12
*/
