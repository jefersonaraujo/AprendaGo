package main

import "fmt"

func main() {

	//declaracao
    x := 10
	y := "olá"
	
	fmt.Println(x,y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	//atribuicao
	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	x, z := 21, 50
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

}