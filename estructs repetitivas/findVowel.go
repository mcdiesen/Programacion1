//Buscar y escribir la primera vocal del teclado(leer uno a uno los caracteres)
package main

import "fmt"

func main() {
	var vocal int
	fmt.Println("Ingrese una vocal:")
	fmt.Scanf("%c", &vocal)
	switch vocal {
	case 'a':
		{
			fmt.Printf("vocal ingresada: %c\n", vocal)
			break
		}
	case 'e':
		{
			fmt.Printf("vocal ingresada: %c\n", vocal)
			break
		}
	case 'i':
		{
			fmt.Printf("vocal ingresada: %c\n", vocal)
			break
		}
	case 'o':
		{
			fmt.Printf("vocal ingresada: %c\n", vocal)
			break
		}
	case 'u':
		{
			fmt.Printf("vocal ingresada: %c\n", vocal)
			break
		}
	default:
		fmt.Println("ninguna vocal ingresada")
		break
	}
}
