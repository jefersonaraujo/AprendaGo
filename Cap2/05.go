package main

import "fmt"

func main() {

	x := `"oi bom dia\ncomo vai?\tespero \"que\" tudo bem."`
	y := "oi bom dia\ncomo vai?\tespero \"que\" tudo bem."
	fmt.Println(x)
	fmt.Println(y)

	z := "oi"
	w := "bom dia"
	u := fmt.Sprint(z,w)

	
	fmt.Println(u)


	 

}
 