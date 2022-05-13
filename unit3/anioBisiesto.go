package main

import (
	"fmt"
)

func esBisiesto(aa int) bool {
	return aa%4 == 0 && aa%100 != 0 || aa%400 == 0
}
func validarFecha(dd, mm, aa int) bool {
	var Valid bool
	if dd == 0 || mm == 0 || aa == 0 {
		Valid = true
	} else if (dd < 1) || (dd > 31) {
		Valid = true
	}
	switch mm {
	case 2:
		if dd > 29 {
			Valid = true
		}
	case 4, 6, 9, 11:
		if dd > 30 {
			Valid = true
		}
	}
	return Valid
}
func main() {
	var dd, mm, aa int
	fmt.Println("Ingrese el mes y el año: ")
	fmt.Scan(&dd, &mm, &aa)

	validaFecha := validarFecha(dd, mm, aa)

	if validaFecha == true {
		fmt.Println("El día o el mes no es válido")
	}
	bisiesto := esBisiesto(aa)
	if bisiesto == true {
		fmt.Println("Es un año bisiesto", bisiesto)
	} else {
		fmt.Println("No es un año bisiesto")
	}
}
