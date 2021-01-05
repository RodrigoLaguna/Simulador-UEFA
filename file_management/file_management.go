package file_management

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Esta Funcion nos ayuda con el archivo de los equiṕos para separarlos,
//tomando en cuenta un asterisco como delimitador
func Read_E(archivo string) []string {
	//El contenido del archivo se convierte a un arreglo de bytes
	bytesLeidos, err := ioutil.ReadFile(archivo)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	//Convierte los bytes leidos en string
	texto_contenido := string(bytesLeidos)
	//Guardamos cada string en una posicion del arreglo
	s := strings.Split(texto_contenido, "*")
	return s
}

//Funcion que separa los datos de cada equipo/jugador
func D_equipo(equipo []string, posicion int) []string {
	datos := strings.Fields(equipo[posicion])
	return datos
}

//Funcion que va a separar a los jugadores a partir de un salto de linea
func C_jugador(equipo []string, posicion int) []string {
	datos := strings.Split(equipo[posicion], "\n")
	return datos
}

func D_jugador(equipo []string, posicion int) []string {
	datos := strings.Fields(equipo[posicion])
	return datos
}

//Funcion que lee un archivo y devuelve su contenido en arreglo de strings
func Read(archivo string) []string {
	var contenido []string
	//El contenido del archivo se convierte a un arreglo de bytes
	bytesLeidos, err := ioutil.ReadFile(archivo)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	//Convierte los bytes leidos en string
	texto_contenido := string(bytesLeidos)
	//Guardamos cada string en una posicion del arreglo
	contenido = strings.Fields(texto_contenido)
	return contenido
}
