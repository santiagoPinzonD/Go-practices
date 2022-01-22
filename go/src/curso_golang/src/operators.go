package main

import "fmt"

func main() {
	x := 20
	y := 10

	// suma 
	result := x + y
	fmt.Println(result)

	// resta
	result = y - x
	fmt.Println(result)

	// Multiplicacion
	result = y * x
	fmt.Println(result)

	// Division
	result = y / x
	fmt.Println(result)

	// Modulo
	result = x % 10
	fmt.Println(result)

	// Incremental
	x++
	fmt.Println(x)

	// Incremental
	x--
	fmt.Println(x)

	// Retos 
	base := 14
	altura := 7
	areaRectangulo := base * altura
	fmt.Println(areaRectangulo)
}
