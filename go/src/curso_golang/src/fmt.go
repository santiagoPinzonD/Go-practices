package main

import "fmt"

func main() {

	helloMessage := "Hello"
	worlMessage := "World"

	//Println

	fmt.Println(helloMessage, worlMessage)

	//Printf
	name := "santiago"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos \n", name, cursos)
	fmt.Printf("%v tiene mas de %v cursos \n", name, cursos)

	// Sprintf

	message := fmt.Sprintf("%s tiene mas de %d cursos", name, cursos)
	fmt.Println(message)

	// Type of data

	fmt.Printf("helloMessage: %T\n", helloMessage)
}
