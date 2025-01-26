package main

import "fmt"

func main() {
	var edad int = 18

	if edad >= 65 {
		fmt.Println("La Persona esta jubilada")
	} else if edad >= 18 {
		fmt.Println("La Persona esta economicamente activa")
	} else {
		fmt.Println("La persona es menor de edad")
	}
	fmt.Println("Fin del Programa")
}
