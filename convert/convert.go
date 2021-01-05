package convert

import (
	"fmt"
	"strconv"
)

/*
	Función que recibe una array de strings y un valor entero para ubicar la posicion
	a la que se desea convertir a entero

*/
func String_Int(equipos []string, posición int) int {
	//Se crean dos variables de tipo string
	entero, error := strconv.Atoi(equipos[posición])
	if error != nil {
		panic(error)
	}
	return entero
}

/*
	Función que recibe una array de strings y un valor para ubicar la posicion
	para convertir esa posicion en un tipo string
*/
func C_string(cadena []string, reco_equipo int) string {
	tipo := fmt.Sprintf("%s", cadena[reco_equipo])
	return tipo
}
