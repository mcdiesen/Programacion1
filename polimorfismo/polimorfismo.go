//Polimorgfismo
//Patron de diseño Factory Method
package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {
}

func (p Paypal) Pay() {
	fmt.Println("Pagado con Paypal")
}

type Cash struct {
}

func (c Cash) Pay() {
	fmt.Println("Pagado en efectivo")
}

type CreditCard struct {
}

func (cr CreditCard) Pay() {
	fmt.Println("Pagado con Tarjeta de credito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}
func main() {
	var method uint
	fmt.Println("Digite un mètodo de metodo de pago")
	fmt.Println("\n 1. Paypal \n 2. Efectivo\n 3. Tarjeta de crétido")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("debe digitar un metodo valido")
	}
	if method > 3 {
		panic("debe digital un numero valido")
	}
	payMethod := Factory(method)
	payMethod.Pay()
}
