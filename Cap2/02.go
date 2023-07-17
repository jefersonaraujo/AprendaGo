package main

import "fmt"

var y = 11 // package level scope
func main() {
	z := 10
	qualquercoisa(z)
}

func qualquercoisa (x int) {
	fmt.Println(y)
	fmt.Println(x)
}