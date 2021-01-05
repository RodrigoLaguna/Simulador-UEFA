package main

import (
	"Segundo_Semestre/Proyecto_FInal/convert"
	"Segundo_Semestre/Proyecto_FInal/file_management"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type equipos struct {
	numero           int
	gol              int
	nombre           string
	liga             string
	estadio          string
	capacidad        int
	plantillaTitular [11]jugadores
	estadistica      estadisticas
	izquierda        *equipos
	derecha          *equipos
}

type jugadores struct {
	nombre        string
	numero        int
	nacionalidad  string
	posicion      string
	golesAnotados int
}

type estadisticas struct {
	jugados         int
	ganados         int
	perdidos        int
	golesFavor      int8
	golesContra     int8
	diferenciaGoles int8
}

var (
	raiz *equipos
	//Variable que contendra la ruta del archivo
	club   = "plantillas.txt"
	player = "jugadores.txt"

	//Se crea un fichero para la escritura de resultados
	fichero, err = os.Create("resultados.txt")

	//Esta variable recibe un arreglo de strings,la separacion de cada arreglo es un *
	arrayEquipo  = file_management.Read_E(club)
	arrayJugador = file_management.Read_E(player)

	//Variables que sirven para extraer los datos de cada equipo(recorrido)
	recoEquipo int
	j          int
	posicion   int
	cuFinal    [8]int
	semFinal   [4]int
	final      [2]int
)

func main() {

	//Condicion para el manejo de error del archivo a generar
	if err != nil {
		fmt.Print(err)
		return
	}

	//Insertamos el nodo raiz vacio,que servira como comodin al momento de eliminar un nodo
	insertaRaiz()

	fmt.Print("\n\n\t        UEFA Champions League \n")
	fmt.Print("\t     ┏┓ ┏┓ ┏━━━━┓ ┏━━━━┓ ┏━━━━━┓ \n")
	fmt.Print("\t     ┃┃ ┃┃ ┃┃━━━┛ ┃┃━━━┛ ┃┃━━━┃┃ \n")
	fmt.Print("\t     ┃┃ ┃┃ ┃┃━━━┓ ┃┃━━━┓ ┃┃   ┃┃ \n")
	fmt.Print("\t     ┃┃ ┃┃ ┃┃━━━┛ ┃┃━━━┛ ┃┃━━━┃┃ \n")
	fmt.Print("\t     ┃┗━┛┃ ┃┃━━━┓ ┃┃     ┃┃━━━┃┃ \n")
	fmt.Print("\t     ┗━━━┛ ┗━━━━┛ ┗┛     ┗┛   ┗┛ \n\n")

	fmt.Print("\tEste programa realiza una simulacion del\n")
	fmt.Print("\t   campeonato UEFA Liga de Campeones\n\n")
	fmt.Print("\tEl programa se estructura en partidos de\n")
	fmt.Print("\t  eliminacion directa,por lo que cada \n")
	fmt.Print("\t     vez que se elimine un equipo,\n")
	fmt.Print("\t  se escribiran sus datos dentro de \n")
	fmt.Print("\t    un archivo llamado resultados\n\n")

	pausa()

	for i := 1; i < len(arrayJugador)+1; i++ {
		insertar(i)
		recoEquipo++
	}
	fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")
	fmt.Print("\n\n\t          Octavos de Final\n")
	imprimearbol(raiz)
	fmt.Print("\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	pausa()
	for i := 0; i < 8; i++ {
		octavos(raiz)
	}

	fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")
	fmt.Print("\n\n\t         Cuartos de Final\n")
	imprimearbol(raiz)
	fmt.Print("\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	pausa()
	posicion = 0
	j = 0
	for i := 0; i < 4; i++ {
		cuartos(raiz)
	}

	fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")
	fmt.Print("\n\n\t            Semifinal \n")
	imprimearbol(raiz)
	fmt.Print("\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	pausa()
	posicion = 0
	j = 0
	for i := 0; i < 2; i++ {
		semifinal(raiz)
	}

	fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")
	fmt.Print("\n\n\t             Final \n")
	imprimearbol(raiz)
	fmt.Print("   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	pausa()
	posicion = 0
	finalCampeon(raiz)
	fmt.Print("\n\n\t\t¡Fin del torneo!\n\n")
	pausa()
	//Liberarmos la memoria
	elimina(raiz)
	//Cerrar el fichero creado
	fichero.Close()
}

func insertar(x int) {
	//Creacion del nodo
	//Variable que va a separar los datos de cada equipo
	s := file_management.D_equipo(arrayEquipo, recoEquipo)
	linea := file_management.C_jugador(arrayJugador, recoEquipo)
	var nuevo *equipos
	nuevo = new(equipos)
	nuevo.numero = x
	nuevo.nombre = convert.C_string(s, 0)
	nuevo.liga = convert.C_string(s, 1)
	nuevo.estadio = convert.C_string(s, 2)
	nuevo.capacidad = convert.String_Int(s, 3)

	//Datos de sus jugadores titulares
	for i := 0; i < 11; i++ {
		//Variable que va a separar los elementos de cada linea
		elementos := file_management.D_jugador(linea, i+1)
		nuevo.plantillaTitular[i].nombre = convert.C_string(elementos, 0)
		nuevo.plantillaTitular[i].numero = convert.String_Int(elementos, 1)
		nuevo.plantillaTitular[i].nacionalidad = convert.C_string(elementos, 2)
		nuevo.plantillaTitular[i].posicion = convert.C_string(elementos, 3)
		nuevo.plantillaTitular[i].golesAnotados = 0
	}
	nuevo.izquierda = nil
	nuevo.derecha = nil
	//El primer nodo que se inserta es el nodo raiz
	if raiz == nil {
		raiz = nuevo
		//Se crean los demas nodos dependiendo su valor
	} else {
		var anterior, recorrer *equipos
		anterior = nil
		recorrer = raiz
		//Recorremos el arbol
		for recorrer != nil {
			anterior = recorrer
			if x < recorrer.numero {
				recorrer = recorrer.izquierda
			} else {
				recorrer = recorrer.derecha
			}
		}
		//Si el valor es menor a la raiz se inserta a la izquierda
		if x < anterior.numero {
			anterior.izquierda = nuevo
			//Si el valor es mayor a la raiz se inserta a la derecha
		} else {
			anterior.derecha = nuevo
		}
	}

}

func octavos(nodo *equipos) {

	numLoc := j + 1
	rivalUno := busquedaRec(nodo, numLoc)
	numVis := j + 2
	rivalDos := busquedaRec(nodo, numVis)

	fmt.Printf("\n\n\tEstadio: %s           \n", rivalUno.estadio)
	fmt.Printf("\n\tPartido: %s vs %s  \n", rivalUno.nombre, rivalDos.nombre)
	fmt.Println("\n\t\t¡Inicia el partido!")
	//time.Sleep(2 * time.Second)
	fmt.Print("\n\n\t     Informacion del partido\n")
	for i := 0; i < 5; i++ {
		golLoc := valoRand()
		golVis := valoRand()

		//Simulacion de goles y tiros cercanos
		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Gol de %s!(%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			//Funcion que llena las estadisticas
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalUno.nombre)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			rivalDos.plantillaTitular[golJug].golesAnotados++
			rivalDos.estadistica.golesFavor++
			rivalUno.estadistica.golesContra++
			rivalDos.gol++
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalDos.nombre)
			time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Tiempos extra!")
		time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalDos.nombre)
			time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Penales!")
		time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		//Condicion para que se de un ganador
		if golLoc == golVis {
			if golLoc == 1 {
				golLoc = 0
				golVis = 1
			} else {
				golLoc = 1
				golVis = 0
			}
		}

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\n\t¡Fallo Penal! ", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			golJug := valorJug()
			fmt.Println("\n\t¡Fallo Penal! ", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre, "(", rivalDos.nombre, ")")
			time.Sleep(2 * time.Second)
		}

	}

	fmt.Print("\n\n\n\t¡Final del partido!\n\n")
	time.Sleep(2 * time.Second)
	pausa()

	if rivalUno.gol > rivalDos.gol {
		//Antes de eliminar llenamos los datos del ganador y perdedor
		aumentaPartidos(rivalUno, rivalDos)
		//Variable para guardar las posiciones para la siguiente fase
		cuFinal[posicion] = rivalUno.numero
		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalDos)
		escribeJug(rivalDos)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t  Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		eliminaNodo(raiz, rivalDos.numero)

	} else {
		//Antes de eliminar el nodo perdedor llenamos los datos del ganador y perdedor
		aumentaPartidos(rivalDos, rivalUno)
		//Variable para guardar las posiciones para la siguiente fase
		cuFinal[posicion] = rivalDos.numero

		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalUno)
		escribeJug(rivalUno)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		eliminaNodo(raiz, rivalUno.numero)

	}

	//Controlamos las posteriores iteraciones
	j = j + 2
	rivalUno.gol = 0
	rivalDos.gol = 0
	posicion++

	//Randoms para generar estadistica
	faltaLoc := valoFal()
	tAmarillaLoc := valoAma()
	tRojaLoc := valoRoj()
	pos := valoPos()
	faltaVis := valoFal()
	tAmarillaVis := valoAma()
	tRojaVis := valoRoj()

	fmt.Print("\t   **Estadosticas del partido**\n")
	fmt.Printf("\t   %d%c  Posecion del balon %d%c \n", pos, 37, 100-pos, 37)
	fmt.Printf("\t    %d        Faltas        %d \n", faltaLoc, faltaVis)
	fmt.Printf("\t    %d       T.amarillas    %d \n", tAmarillaLoc, tAmarillaVis)
	fmt.Printf("\t    %d        T.Rojas       %d \n\n", tRojaLoc, tRojaVis)
	fmt.Print("   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	pausa()

}

func cuartos(nodo *equipos) {

	numLoc := cuFinal[posicion]
	rivalUno := busquedaRec(nodo, numLoc)
	numVis := cuFinal[posicion+1]
	rivalDos := busquedaRec(nodo, numVis)

	fmt.Printf("\n\n\tEstadio: %s           \n", rivalUno.estadio)
	fmt.Printf("\n\tPartido: %s vs %s  \n", rivalUno.nombre, rivalDos.nombre)
	fmt.Println("\n\t\t¡Inicia el partido!")
	time.Sleep(2 * time.Second)
	fmt.Print("\n\n\t     Informacion del partido\n")
	for i := 0; i < 5; i++ {
		golLoc := valoRand()
		golVis := valoRand()

		//Simulacion de goles y tiros cercanos
		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Gol de %s!(%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			//time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalUno.nombre)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			//time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalDos.nombre)
			//time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Tiempos extra!")
		//time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalDos.nombre)
			time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Penales!")
		time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		//Condicion para que se de un ganador
		if golLoc == golVis {
			if golLoc == 1 {
				golLoc = 0
				golVis = 1
			} else {
				golLoc = 1
				golVis = 0
			}
		}

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\n\t¡Fallo Penal! ", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			golJug := valorJug()
			fmt.Println("\n\t¡Fallo Penal! ", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre, "(", rivalDos.nombre, ")")
			time.Sleep(2 * time.Second)
		}

	}

	fmt.Print("\n\n\n\t¡Final del partido!\n\n")
	time.Sleep(2 * time.Second)
	//pausa()

	if rivalUno.gol > rivalDos.gol {
		//Antes de eliminar el nodo perdedor llenamos los datos del ganador y perdedor
		aumentaPartidos(rivalUno, rivalDos)
		//Variable para guardar las posiciones para la siguiente fase
		semFinal[j] = rivalUno.numero

		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalDos)
		escribeJug(rivalUno)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t  Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		eliminaNodo(raiz, rivalDos.numero)

	} else {
		//Antes de eliminar el nodo perdedor llenamos los datos del ganador y perdedor
		aumentaPartidos(rivalDos, rivalUno)
		//Variable para guardar las posiciones para la siguiente fase
		semFinal[j] = rivalDos.numero

		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalUno)
		escribeJug(rivalUno)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		eliminaNodo(raiz, rivalUno.numero)

	}

	//Controlamos las posteriores iteraciones
	posicion = posicion + 2
	j++
	rivalUno.gol = 0
	rivalDos.gol = 0

	//Randoms para generar estadistica
	faltaLoc := valoFal()
	tAmarillaLoc := valoAma()
	tRojaLoc := valoRoj()
	pos := valoPos()
	faltaVis := valoFal()
	tAmarillaVis := valoAma()
	tRojaVis := valoRoj()

	fmt.Print("\t   **Estadosticas del partido**\n")
	fmt.Printf("\t   %d%c  Posecion del balon %d%c \n", pos, 37, 100-pos, 37)
	fmt.Printf("\t    %d        Faltas        %d \n", faltaLoc, faltaVis)
	fmt.Printf("\t    %d       T.amarillas    %d \n", tAmarillaLoc, tAmarillaVis)
	fmt.Printf("\t    %d        T.Rojas       %d \n\n", tRojaLoc, tRojaVis)
	fmt.Print("   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	pausa()

}

func semifinal(nodo *equipos) {

	numLoc := semFinal[posicion]
	rivalUno := busquedaRec(nodo, numLoc)
	numVis := semFinal[posicion+1]
	rivalDos := busquedaRec(nodo, numVis)

	fmt.Printf("\n\n\tEstadio: %s           \n", rivalUno.estadio)
	fmt.Printf("\n\tPartido: %s vs %s  \n", rivalUno.nombre, rivalDos.nombre)
	fmt.Println("\n\t\t¡Inicia el partido!")
	time.Sleep(2 * time.Second)
	fmt.Print("\n\n\t     Informacion del partido\n")
	for i := 0; i < 5; i++ {
		golLoc := valoRand()
		golVis := valoRand()

		//Simulacion de goles y tiros cercanos
		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Gol de %s!(%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalUno.nombre)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalDos.nombre)
			time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Tiempos extra!")
		time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalDos.nombre)
			time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Penales!")
		time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		//Condicion para que se de un ganador
		if golLoc == golVis {
			if golLoc == 1 {
				golLoc = 0
				golVis = 1
			} else {
				golLoc = 1
				golVis = 0
			}
		}

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\n\t¡Fallo Penal! ", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			golJug := valorJug()
			fmt.Println("\n\t¡Fallo Penal! ", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre, "(", rivalDos.nombre, ")")
			time.Sleep(2 * time.Second)
		}

	}

	fmt.Print("\n\n\n\t¡Final del partido!\n\n")
	time.Sleep(2 * time.Second)
	//pausa()
	//

	if rivalUno.gol > rivalDos.gol {
		//Antes de eliminar el nodo perdedor llenamos los datos del ganador y perdedor
		aumentaPartidos(rivalUno, rivalDos)
		//Variable para guardar las posiciones para la siguiente fase
		final[j] = rivalUno.numero

		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalDos)
		escribeJug(rivalDos)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t  Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		eliminaNodo(raiz, rivalDos.numero)

	} else {
		//Antes de eliminar el nodo perdedor llenamos los datos del ganador y perdedor
		aumentaPartidos(rivalDos, rivalUno)
		//Variable para guardar las posiciones para la siguiente fase
		final[j] = rivalDos.numero

		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalUno)
		escribeJug(rivalUno)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		eliminaNodo(raiz, rivalUno.numero)

	}

	//Controlamos las posteriores iteraciones
	posicion = posicion + 2
	j++
	rivalUno.gol = 0
	rivalDos.gol = 0

	//Randoms para generar estadistica
	faltaLoc := valoFal()
	tAmarillaLoc := valoAma()
	tRojaLoc := valoRoj()
	pos := valoPos()
	faltaVis := valoFal()
	tAmarillaVis := valoAma()
	tRojaVis := valoRoj()

	fmt.Print("\t   **Estadosticas del partido**\n")
	fmt.Printf("\t   %d%c  Posecion del balon %d%c \n", pos, 37, 100-pos, 37)
	fmt.Printf("\t    %d        Faltas        %d \n", faltaLoc, faltaVis)
	fmt.Printf("\t    %d       T.amarillas    %d \n", tAmarillaLoc, tAmarillaVis)
	fmt.Printf("\t    %d        T.Rojas       %d \n\n", tRojaLoc, tRojaVis)
	fmt.Print("   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	pausa()

}

func finalCampeon(nodo *equipos) {

	numLoc := final[posicion]
	rivalUno := busquedaRec(nodo, numLoc)
	numVis := final[posicion+1]
	rivalDos := busquedaRec(nodo, numVis)

	fmt.Printf("\n\n\tEstadio: %s           \n", rivalUno.estadio)
	fmt.Printf("\n\tPartido: %s vs %s  \n", rivalUno.nombre, rivalDos.nombre)
	fmt.Println("\n\t\t¡Inicia el partido!")
	//time.Sleep(2 * time.Second)
	fmt.Print("\n\n\t     Informacion del partido\n")
	for i := 0; i < 5; i++ {
		golLoc := valoRand()
		golVis := valoRand()

		//Simulacion de goles y tiros cercanos
		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Gol de %s!(%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalUno.nombre)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("\n\tTiro del %s cercano a porteria \n ", rivalDos.nombre)
			time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Tiempos extra!")
		time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\tGol de %s (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\tTiro cercano a porteria por parte del", rivalDos.nombre)
			time.Sleep(2 * time.Second)
		}
	}

	if rivalUno.gol == rivalDos.gol {
		fmt.Println("\t\t¡Penales!")
		//time.Sleep(2 * time.Second)
		golLoc := valoRand()
		golVis := valoRand()

		//Condicion para que se de un ganador
		if golLoc == golVis {
			if golLoc == 1 {
				golLoc = 0
				golVis = 1
			} else {
				golLoc = 1
				golVis = 0
			}
		}

		if golLoc == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalUno.plantillaTitular[golJug].nombre, rivalUno.nombre)
			aumentaGol(rivalUno, rivalDos, golJug)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("\n\t¡Fallo Penal! ", rivalUno.nombre)
			time.Sleep(2 * time.Second)
		}

		if golVis == 1 {
			golJug := valorJug()
			fmt.Printf("\n\t¡Anota penal %s! (%s)\n", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre)
			aumentaGol(rivalDos, rivalUno, golJug)
			time.Sleep(2 * time.Second)
		} else {
			golJug := valorJug()
			fmt.Println("\n\t¡Fallo Penal! ", rivalDos.plantillaTitular[golJug].nombre, rivalDos.nombre, "(", rivalDos.nombre, ")")
			time.Sleep(2 * time.Second)
		}

	}

	fmt.Print("\n\n\n\t¡Final del partido!\n\n")
	time.Sleep(2 * time.Second)
	//pausa()

	if rivalUno.gol > rivalDos.gol {
		//Antes de eliminar el nodo perdedor llenamos los datos del ganador y perdedor
		rivalUno.estadistica.jugados++
		rivalUno.estadistica.ganados++
		rivalDos.estadistica.jugados++
		rivalDos.estadistica.perdidos++
		rivalDos.estadistica.diferenciaGoles = rivalDos.estadistica.golesFavor - rivalDos.estadistica.golesContra
		rivalUno.estadistica.diferenciaGoles = rivalUno.estadistica.golesFavor - rivalUno.estadistica.golesContra

		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalDos)
		escribeJug(rivalDos)
		escribeCampeon(rivalUno)
		escribeJug(rivalUno)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t  Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		fmt.Printf("\n\n\t    ¡%s Campeon!\n\n", rivalUno.nombre)
		eliminaNodo(raiz, rivalDos.numero)

	} else {
		//Antes de eliminar el nodo perdedor llenamos los datos del ganador y perdedor
		rivalDos.estadistica.jugados++
		rivalDos.estadistica.ganados++
		rivalUno.estadistica.jugados++
		rivalUno.estadistica.perdidos++
		rivalUno.estadistica.diferenciaGoles = rivalUno.estadistica.golesFavor - rivalUno.estadistica.golesContra
		rivalDos.estadistica.diferenciaGoles = rivalDos.estadistica.golesFavor - rivalDos.estadistica.golesContra

		//Escribimos los datos en el archivo txt
		escribeEliminado(rivalUno)
		escribeJug(rivalUno)
		escribeCampeon(rivalDos)
		escribeJug(rivalDos)
		fmt.Print("\n\n   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
		fmt.Print("\t\t Marcador Final\n\n")
		fmt.Printf("\t%s [%d] vs [%d] %s\n\n", rivalUno.nombre, rivalUno.gol, rivalDos.gol, rivalDos.nombre)
		fmt.Printf("\n\n\t    ¡%s Campeon!\n\n", rivalDos.nombre)
		eliminaNodo(raiz, rivalUno.numero)

	}

	//Randoms para generar estadistica
	faltaLoc := valoFal()
	tAmarillaLoc := valoAma()
	tRojaLoc := valoRoj()
	pos := valoPos()
	faltaVis := valoFal()
	tAmarillaVis := valoAma()
	tRojaVis := valoRoj()

	fmt.Print("\t   **Estadosticas del partido**\n")
	fmt.Printf("\t   %d%c  Posecion del balon %d%c \n", pos, 37, 100-pos, 37)
	fmt.Printf("\t    %d        Faltas        %d \n", faltaLoc, faltaVis)
	fmt.Printf("\t    %d       T.amarillas    %d \n", tAmarillaLoc, tAmarillaVis)
	fmt.Printf("\t    %d        T.Rojas       %d \n\n", tRojaLoc, tRojaVis)
	fmt.Print("   °°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
}

//Funcion que ayuda a buscar un equipo para posterior eliminarlo
func busquedaRec(arbol *equipos, dato int) *equipos {
	r := arbol
	if arbol == nil {
		return r
	}
	if dato < arbol.numero {
		r = busquedaRec(arbol.izquierda, dato)
	} else if dato > arbol.numero {
		r = busquedaRec(arbol.derecha, dato)
	} else {
		r = arbol // son iguales, lo encontre
	}
	return r
}

///////////////////////Funciones Random/////////////////////

//Funcion principal para el uso de random
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

//Funcion para dar random a los goles
func valoRand() int {
	rand.Seed(time.Now().UnixNano())
	num := random(0, 2)
	return num
}

//Funcion para dar random a tarjetas amarillas
func valoAma() int {
	rand.Seed(time.Now().UnixNano())
	num := random(0, 4)
	return num
}

//Funcion para dar random a tarjetas rojas
func valoRoj() int {
	rand.Seed(time.Now().UnixNano())
	num := random(0, 3)
	return num
}

//Funcion para dar random a la poseción del balón
func valoPos() int {
	rand.Seed(time.Now().UnixNano())
	num := random(0, 100)
	return num
}

//Funcion para dar random a las faltas cometidas
func valoFal() int {
	rand.Seed(time.Now().UnixNano())
	num := random(0, 10)
	return num
}

//Funcion para aignar los goles anotados a un jugador random
func valorJug() int {
	rand.Seed(time.Now().UnixNano())
	num := random(0, 11)
	return num
}

/////////////////////////////////////////////////////////////

//Esta funcion elimina un nodo y reconecta los restantes
func eliminaNodo(root *equipos, data int) *equipos {
	if root == nil {
		return root
	} else if data < root.numero {
		root.izquierda = eliminaNodo(root.izquierda, data)
	} else if data > root.numero {
		root.derecha = eliminaNodo(root.derecha, data)
	} else {
		if root.izquierda == nil && root.derecha == nil { //nodo que no tiene hijos
			root = nil
		} else if root.izquierda == nil { //Tiene un hijo
			temp := root
			root = root.derecha
			temp = nil
			if temp == nil {

			}
		} else if root.derecha == nil { //Tiene un hijo
			temp := root
			root = root.derecha
			temp = nil
			if temp == nil {

			}
		} else { //Tiene dos hijos
			temp := root.derecha

			for temp.izquierda != nil {
				temp = temp.izquierda
			}
			root.numero = temp.numero
			root.derecha = eliminaNodo(root.derecha, temp.numero)
		}
	}
	return root
}

//Funcion que ingresa el nodo raiz vacio,esto facilita la eliminacion de un nodo
func insertaRaiz() {
	var nuevo *equipos
	nuevo = new(equipos)
	nuevo.numero = 0
	nuevo.izquierda = nil
	nuevo.derecha = nil
	raiz = nuevo
}

//Función que pausa el programa
func pausa() {
	fmt.Print("\tPresiona 'Enter' para continuar ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

//Función que escribe los datos del equipo eliminado
func escribeEliminado(nodo *equipos) {
	fmt.Fprint(fichero, "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	fmt.Fprintln(fichero, nodo.nombre)
	fmt.Fprintln(fichero, " ", nodo.estadistica.jugados, "\tJuegos jugados")
	fmt.Fprintln(fichero, " ", nodo.estadistica.ganados, "\tVictorias")
	fmt.Fprintln(fichero, " ", nodo.estadistica.ganados, "\tVictorias")
	fmt.Fprintln(fichero, " ", nodo.estadistica.perdidos, "\tDerrotas")
	fmt.Fprintln(fichero, " ", nodo.estadistica.golesFavor, "\tGoles anotados")
	fmt.Fprintln(fichero, " ", nodo.estadistica.golesContra, "\tGoles en contra")
	fmt.Fprintln(fichero, " ", nodo.estadistica.diferenciaGoles, "\tDif.Goles")
	fmt.Fprint(fichero, "\n\n")
}

//Imprime el arbol
func imprimearbol(p *equipos) {
	if p == nil {
		return
	} else {
		fmt.Printf("\t\t  %s \n", p.nombre)
		imprimearbol(p.izquierda)
		imprimearbol(p.derecha)
	}
}

//Liberar memoria la final del programa
func elimina(p *equipos) {
	if p != nil {
		elimina(p.izquierda)
		elimina(p.derecha)
		p = nil
		raiz = p
	}
}

//Función que escribe en el archivo de texto los datos del campeón
func escribeCampeon(nodo *equipos) {
	fmt.Fprint(fichero, "°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n\n")
	fmt.Fprintln(fichero, "\tCampeon")
	fmt.Fprintln(fichero, nodo.nombre)
	fmt.Fprintln(fichero, " ", nodo.estadistica.jugados, "\tJuegos jugados")
	fmt.Fprintln(fichero, " ", nodo.estadistica.ganados, "\tVictorias")
	fmt.Fprintln(fichero, " ", nodo.estadistica.ganados, "\tVictorias")
	fmt.Fprintln(fichero, " ", nodo.estadistica.perdidos, "\tDerrotas")
	fmt.Fprintln(fichero, " ", nodo.estadistica.golesFavor, "\tGoles anotados")
	fmt.Fprintln(fichero, " ", nodo.estadistica.golesContra, "\tGoles en contra")
	fmt.Fprintln(fichero, " ", nodo.estadistica.diferenciaGoles, "\tDif.Goles")
	fmt.Fprint(fichero, "\n\n")
}

//Función que escribe en el archivo de texto los datos de jugadores
func escribeJug(nodo *equipos) {
	fmt.Fprintln(fichero, "Jugadores")
	fmt.Fprintln(fichero, "Goles \tNombre")
	for i := 0; i < 11; i++ {
		fmt.Fprint(fichero, "  ", nodo.plantillaTitular[i].golesAnotados)
		fmt.Fprint(fichero, "\t", nodo.plantillaTitular[i].nombre)
		fmt.Fprint(fichero, "\n")
	}
	fmt.Fprint(fichero, "\n\n")
}

//Funcion que aumenta la estadistica de goles
func aumentaGol(nodo *equipos, nodoDos *equipos, jugador int) {
	nodo.plantillaTitular[jugador].golesAnotados++
	nodo.estadistica.golesFavor++
	nodoDos.estadistica.golesContra++
	nodo.gol++
}

//Función que aumenta la estadistica de partidos
func aumentaPartidos(nodo *equipos, nodoDos *equipos) {
	nodo.estadistica.jugados++
	nodo.estadistica.ganados++
	nodoDos.estadistica.jugados++
	nodoDos.estadistica.perdidos++
	nodoDos.estadistica.diferenciaGoles = nodoDos.estadistica.golesFavor - nodoDos.estadistica.golesContra
}
