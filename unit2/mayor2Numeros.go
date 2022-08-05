package main

//obtener el mayor de dos números ingresados a través del teclado
import "fmt"

func main() {
	var numero1, numero2 int

	fmt.Println("Ingrese el número entero 1:")
	fmt.Scanf("%d", &numero1)

	fmt.Println("Ingrese el número entero 2:")
	fmt.Scanf("%d", &numero2)

	if numero1 > numero2 {
		fmt.Println("Número", numero1, "es mayor que", numero2)
	} else {
		fmt.Println("Número", numero2, "es mayor que", numero1)
	}
}
