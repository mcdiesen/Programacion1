package main

import "fmt"

func main(){
var A, B int
A=2
B=5
fmt.Println("Ingrese los valores para A  y B respectivamente:")
fmt.Scan(&A)
fmt.Scan(&B)

resultado:= 3*A-4*B/A^2
fmt.Println("El resultdo es: ",resultado)
}
