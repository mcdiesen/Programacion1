//Calcula la suma y media de 20 elemantos de un arreglo
package main

import "fmt"

func main() {
	var Suma, media float32
	var arreglo1 [20]int

	for _, elem := range arreglo1 {
		fmt.Scanf("%d", &arreglo1[elem])
		Suma = Suma + float32(arreglo1[elem])
	}
	media = Suma / 2.0
	fmt.Println("Suma", Suma)
	fmt.Printf("Media %.1f", media)
}
