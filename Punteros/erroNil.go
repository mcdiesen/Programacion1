//puntero nil
package main

import "fmt"

func main() {
	first := 100
	var second *int
	//second = &first
	fmt.Println(first, second)
	fmt.Println(second == nil)
}
