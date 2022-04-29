//enums(enumerados) es un tipo de dato que consiste en valores constantes
//veamos un ejemplo con los días de la semana
package main

import "fmt"

type DayOfTheWeek uint8

//declaracion de constante
const (
	Monday DayOfTheWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Printf("Monday is %d\n", Monday)
	fmt.Printf("Wednesday is %d\n", Wednesday)
	fmt.Printf("Friday is %d\n", Friday)
}

/*
ejecutar el programa:
go run enumsVar.go
Salida:
	Monday is 0
	Wednesday is 2
	Friday is 4
*/
