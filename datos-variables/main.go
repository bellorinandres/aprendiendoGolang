package main

import "fmt"

func main() {
	/* Parte de Cadenas de Texto o variables */

	// Agregar una variable declarando lo que es, tipo string "Cadena de Texto"
	var nombrePersona string = "Andres"
	// Agregar una variable para que golang la detecte automaticamente
	var apellidoPersona = "Bellorin"
	//Una declaracion de variable mas informal es :=
	segundoNombre := "David"
	/* Concatenar con + */
	fmt.Println("Hola Mundo. Soy " + nombrePersona + " " + segundoNombre + " " + apellidoPersona)

	/* Parte Numerica */

	var anoActual int16 = 2024
	edad := 28
	// Concatenar con Numeros es con ,
	fmt.Println("El año actual es => ", anoActual)
	fmt.Println("Mi Edad => ", edad)
}
