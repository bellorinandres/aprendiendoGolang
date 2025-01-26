package main

import "fmt"

func mapas() {
	// Esto es un mapa, funciona como un arreglo variable valor
	miMapa := map[string]string{
		"Colombia":  "Bogota",
		"Venezuela": "Caracas",
		"Brazil":    "Brasilia",
		"Chile":     "Santiago",
	}
	fmt.Println("El mapa de paises es: ", miMapa)
	fmt.Println("La capital de Venezuela es:", miMapa["Venezuela"])
	// Agregar variables al mapa
	miMapa["Argentina"] = "Buenos Aires"
	fmt.Println("La ciudad de ", miMapa["Argentina"])
	// Borramos un valor del Mapa
	delete(miMapa, "Chile")
}
