package main

import "fmt"

func fechaValida(dd, mm, aa int) bool {
	//	var fechaValida = true

	if (dd == 0) || (mm == 0) || (aa == 0) {
		return false
	}
	if (mm < 1) || (mm > 12) {
		return false
	}
	if dd < 1 {
		return false
	} else {
		switch mm {
		case 4, 6, 9, 11:
			{
				if dd > 30 {
					return false
				}
			}
		case 2:
			{
				if esBisiesto(aa) && (dd > 29) {
					return false
				} else if esBisiesto(aa) && (dd > 28) {
					return false
				} else if dd > 31 {
					return false
				}

			}
		}
	}
	return true
}

func esBisiesto(aa int) bool {
	return (((aa%4 == 0) && (aa%100 != 0)) || (aa%400 == 0))
}
func main() {
	var dd, mm, aa int
	fmt.Println("Ingrese el día, mes y año")
	fmt.Scan(&dd, &mm, &aa)
	valido := fechaValida(dd, mm, aa)
	if valido == true {
		fmt.Println("La fecha es valida", valido)
	} else {
		fmt.Println("La fecha no es válida", valido)
	}
}
