/* Arreglos */

package main

import "fmt"

func main() {
	/* Declaracion de Arreglos */
	// Se puede crear un arreglo sin variables dentro
	// var listaFrutas = [4] string{}
	var listaFrutas = [4]string{"Pera", "Mango", "Fresa"}
	fmt.Println("Esto es un arreglo ",listaFrutas)
	// Imprimir posicion exacta del arreglo
	fmt.Println("Imprimir una posicion exacta del arreglo ",listaFrutas[0])

	/* Slices es crear un arreglo con un tamaño dinamico */

	listaPaises := []string{"Chile", "Peru", "Venezuela"}
	fmt.Println("Imprimir un arreglo por consola ",listaPaises)

	// Agregar un valor al final del arreglo.
	listaPaises = append(listaPaises, "Brasil")
	fmt.Println("Agregamos una lista de paises despues con append ",listaPaises)

	// EL ultimo rango no va incluido.
	listaPaises2 := listaPaises[1:3]
	fmt.Println("El ultimo rango no va incluido ",listaPaises2)

	listaPaises4 := listaPaises[1:]
	fmt.Println("Desde el rango [1] hasta el final ",listaPaises4)

	listaPaises3 := listaPaises[:3]
	fmt.Println("Desde el un rango posterior hasta el inicio ",listaPaises3)
}
