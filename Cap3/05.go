package main

import (
	"fmt"
)

type meutipovar int
var x meutipovar
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	//conversao
	fmt.Println("↑ foi x.\n↓ é y!")
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	

}