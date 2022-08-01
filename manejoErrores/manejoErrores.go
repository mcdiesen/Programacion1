//Control de errores
//si el divisor es 0 controla el error
package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error) {
	//control del error
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre cero!")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {
	result, error := division(4, 2)
	if error == nil {
		fmt.Println("La division es:", result)
	} else {
		fmt.Println("Error no dividir por cero")
	}

}
