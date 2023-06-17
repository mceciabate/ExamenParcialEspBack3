package tickets

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id             string
	NombreCompleto string
	Email          string
	PaisDestino    string
	HoraVuelo      string
	Precio         string
}

// TODO PREGUNTAR POR EJECUCIÓN DEL DEFER
// var archivo string = "tickets.csv"

func manejoPanics() {
	a := recover()
	if a != nil {
		fmt.Println("Cortando ejecución: ", a)
	}
}

// Función para obtener datos a partir del archivo.csv
func ObtenerDatos(ruta string) ([]Ticket, error) {
	var array []Ticket
	// var newTicket Ticket
	// var line4 string
	var arrayStrings [][]string
	// var arrayTicket []string
	defer manejoPanics()
	rawData, er := os.ReadFile(ruta)
	if er != nil {
		panic("Error de lectura de archivo")
	}
	data := strings.Split(string(rawData), "\n ")
	for _, v := range data {
		line := strings.Split(v, "\n")
		for _, v := range line {
			line2 := strings.Split(v, " ")
			arrayStrings = append(arrayStrings, line2)
		}
		for i := 0; i < len(arrayStrings); i++ {
			line3 := arrayStrings[i]
			line4 := strings.Join(line3, " ")
			arrayTicket := strings.Split(line4, ",")
			newTicket := Ticket{
				Id:             arrayTicket[0],
				NombreCompleto: arrayTicket[1],
				Email:          arrayTicket[2],
				PaisDestino:    arrayTicket[3],
				HoraVuelo:      arrayTicket[4],
				Precio:         arrayTicket[5],
			}
			array = append(array, newTicket)
		}
		if len(array) == 0 {
			return nil, errors.New("\nNo se ha generado el listado de forma correcta")
		}
	}
	return array, nil
}

// Función para obtener el listado de Tickets según destino
func ObtenerTotalTicketsDestino(destino string, a []Ticket) (int, error) {
	// ListadoRecuperadoTickets, err := ObtenerDatos(archivo)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	acum := 0
	for _, v := range a {
		if v.PaisDestino == destino {
			acum++
		}
	}
	if acum == 0 {
		return 0, errors.New("\nNo se encontraron coincidencias con el destino")
	}
	return acum, nil
}

// Función para obtener Tickets según franja horaria
func ObtenerTicketsFranjaHoraria(time string, a []Ticket) (int, error) {
	var cont int = 0
	for _, v := range a {
		hora := strings.Split(v.HoraVuelo, ":")
		horaInt, err := strconv.Atoi(hora[0])
		if err != nil {
			panic("Error de conversión de dato")
		}
		switch time {
		case "Madrugada":
			if horaInt >= 0 && horaInt <= 6 {
				cont++
			}

		case "Mañana":
			if horaInt >= 7 && horaInt <= 12 {
				cont++
			}

		case "Tarde":
			if horaInt >= 13 && horaInt <= 19 {
				cont++
			}

		case "Noche":
			if horaInt >= 20 && horaInt <= 23 {
				cont++
			}

		}
	}
	if cont != 0 {
		fmt.Printf("\nLa cantidad total de tickets para la %s es %d\n", time, cont)

		return cont, nil
	}
	return 0, errors.New("\nIngrese una franja horaria válida")

}

// Función para obtener porcentaje según destino
func ObtenerPromedioDestinos(destino string, a []Ticket) (float64, error) {
	// ListadoRecuperadoTickets, err := ObtenerDatos(archivo)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// e := errors.New("Error en el listado")
	totalListado := float64(len(a))
	// if totalListado == 0 {
	// 	fmt.Println(e)
	// 	return 0, e
	// }
	totalDestinos, er := ObtenerTotalTicketsDestino(destino, a)
	if er != nil {
		// fmt.Println(er)
		return 0, er
	}
	parseTotalDestinos := float64(totalDestinos)

	porcentaje := (parseTotalDestinos * 100) / totalListado
	fmt.Printf("\nEl porcentaje total de tickets para el destino %s es %.2f", destino, porcentaje)

	return porcentaje, nil
}
