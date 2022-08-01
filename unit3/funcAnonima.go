package main

import "fmt"

func main() {
	//func anonima
	//func(){
	//	fmt.Println("Hola desde la funcion anonima")
	//}()
	myFunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s desde la funcion anonima", nombre)
	}
	fmt.Println(myFunc("Charles"))
}
