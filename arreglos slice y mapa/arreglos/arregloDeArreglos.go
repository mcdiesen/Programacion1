package main

import "fmt"

func main() {
	//Arreglo de string de 8x8
	var board [8][8]string
	//asignación de la cadena r en la posición 00 y en la posición 07
	board[0][0] = "r"
	//board[0][4] = "x"
	board[0][7] = "r"

	for column := range board {
		board[1][column] = "p"
	}
	fmt.Println(board)
}
