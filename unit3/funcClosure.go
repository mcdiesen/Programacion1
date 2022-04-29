//function closure. The function accuulator define a closure function that is bounded to variable i
//Statements a:=accumulator(1) and b:=accumulator(2) create two functions with different starting i variables.

package main

import "fmt"

func accumulator(increment int) func() int {
	i := 0
	return func() int {
		i = i + increment
		return i
	}
}

func main() {
	a := accumulator(1)
	b := accumulator(2)

	fmt.Println("a", "b")
	for i := 0; i < 5; i++ {
		fmt.Println(a(), b())
	}
}
