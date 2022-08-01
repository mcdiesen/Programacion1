package main

import (
	"fmt"
	"os"
)

func main() {
	if file, error := os.Open("hoo.txt"); error != nil {
		panic("No es posible obtener el archivo!!")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		log, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:log])
		fmt.Println(contenidoArchivo)
	}

}
