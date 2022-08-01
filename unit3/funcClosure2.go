package main

import (
	"fmt"
	"strings"
)

//repetir un caracter enviado
//closure
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}
func main() {
	repeat3 := repeat(3)
	repeat3("Charles")
	f`mt.Println("repeat3", val)
	
}
