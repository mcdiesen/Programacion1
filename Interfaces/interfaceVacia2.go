//Interfaces nulas y Interfaces vacias
//Intefaces vacia --> Sirven para manajar tipos desconocidos
package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	X()
}

//Wrapper interfaz vacia, recibe cualquier type como parámetro
func wrapper(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
	//verifica el type de la interfaz == string
	//v, ok := i.(string)
	//if ok {
	//	fmt.Printf("\n", strings.ToUpper(v))
	//}
	switch v := i.(type) {
	case string:
		fmt.Printf("\n", strings.ToUpper(v))
	case int:
		fmt.Println(v, "x 2 =", v*2)
	default:
		fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
	}

}

func main() {
	//	var e exampler
	//	fmt.Printf("Valor: %V, Tipo: %T", e, e)
	wrapper(12)
	wrapper(14.45)
	wrapper("mclean")
	wrapper(false)
	wrapper("charles")
}
