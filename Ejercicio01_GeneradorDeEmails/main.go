package main

import (
	"fmt"
)

func main() {
	s := GenerarMails([]string{"Pedro Perez", "Pablo Perez", "Marta Sosa"})
	fmt.Println(s)
}

/*
	GenerarMails, recibe una lista de nombres formados como "NOMBRE APELLIDO" y devuelve
	una lista de emails manteniendo el orden de la lista de entrada, los emails deben formarse como:
	primera letra del nombre, seguido del apellido, seguido de @ml.com (Pepe Perez = pperez@ml.com)
	IMPORTANTE: los emails deben ser completamente en minuscula y no deberian repetirse,
	para ello se les agrega un numero antes del simbolo '@' 1,2,3,...
*/
func GenerarMails(nombres []string) []string {

	// SOLUCIÓN

	return []string{}
}
