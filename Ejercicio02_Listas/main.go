package main

import (
	"fmt"
)

func main() {
	ListaDePares, ListaDeImpares, ListaDeSuma := OrdenarYSumar([]int{12, 5, 12, 2, 33, 67, 99, 44, 3, 4, 67, 321, 9})
	fmt.Println("Lista De Pares", ListaDePares)
	fmt.Println("Lista De Impares", ListaDeImpares)
	fmt.Println("Lista De Suma", ListaDeSuma)
}

/*
	Dado un slice de numeros se pide
		Extraer los pares e impares a listas diferentes y ordenarlos de forma creciente
		Realizar la suma, posicion a posicion almacenando el resultado en otro slice
		La funcion OrdenarYSumar debe devolver los tres slice, en el siguiente orden: ListaDePares, ListaDeImpares, ListaDeSuma
	Ejemplo
		input: [12,21,33,4] -> output: [4,12], [21,33], [25,35]
		input: [12,21,33,4,1] -> output: [4,14], [1,21,33], [5,35,33]
*/
func OrdenarYSumar(numeros []int) ([]int, []int, []int) {

	// SOLUCIÓN

	return []int{}, []int{}, []int{}
}
