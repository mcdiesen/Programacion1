//Sum two numbers passed by arguments
//Los argumentos de Linea de comandos son una forma común para parametrizar la ejecución
//de programas
//ejemplo: go run sum2Nums.go 2 4
//resultado a mostrar: 2 + 4 = 6
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWith := os.Args
	//strvconv.Atoi(argsWith[1] convierte el primer parámetro de type string a type int
	numA, err := strconv.Atoi(argsWith[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	numB, err := strconv.Atoi(argsWith[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	//Result almacena la suma de los dos argumentos pasados al programa
	result := numA + numB
	fmt.Printf("%d + %d = %d\n", numA, numB, result)
}
