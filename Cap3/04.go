package main

import (
	"fmt"
)

type meutipovar int
var x meutipovar

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	

}