//Pruebas unitarias
package testUnitario

import "testing"

/*func TestSuma(t *testing.T) {
	total := suma(5, 5)
	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}
*/
//Prueba de la funcion Suma a través de tabla
func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		//valores para la prueba. c es el resultado de a y b
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}
	for _, item := range tabla {
		total := suma(item.a, item.b)
		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}

func TestMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{5, 3, 5},
	}
	for _, item := range tabla {
		max := GetMax(item.a, item.b)

		if max != item.c {
			t.Errorf("GetMax incorrecta, tiene %d se esperaba %d", max, item.c)
		}
	}
}
