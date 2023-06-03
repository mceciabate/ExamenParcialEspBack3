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

var archivo string = "tickets.csv"

func manejoPanics() {
	a := recover()
	if a != nil {
		fmt.Println("Cortando ejecución: ", a)
	}
}

// Función para obtener datos a partir del archivo.csv
func ObtenerDatos(ruta string) ([]Ticket, error) {
	var array []Ticket
	var newTicket Ticket
	var line4 string
	var arrayStrings [][]string
	var arrayTicket []string
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
			line4 = strings.Join(line3, " ")
			arrayTicket = strings.Split(line4, ",")
			newTicket = Ticket{
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
			return nil, errors.New("No se ha generado el listado de forma correcta")
		}
	}
	return array, nil
}

// Función para obtener el listado de Tickets según destino
func ObtenerTotalTicketsDestino(destino string) (int, error) {
	ListadoRecuperadoTickets, err := ObtenerDatos(archivo)
	if err != nil {
		fmt.Println(err)
	}
	e := errors.New("No se encontraron coincidencias con el destino")
	acum := 0
	for _, v := range ListadoRecuperadoTickets {
		if v.PaisDestino == destino {
			acum++
		}
	}
	if acum == 0 {
		return 0, e
	}
	return acum, nil
}

// Función para obtener Tickets según franja horaria
func ObtenerTicketsFranjaHoraria(time string) (int, error) {
	defer manejoPanics()
	ListadoRecuperadoTickets, err := ObtenerDatos(archivo)
	if err != nil {
		fmt.Println(err)
	}
	e := errors.New("Ingrese una franja horaria válida")
	var listaMañana []int
	var listaTarde []int
	var listaNoche []int
	var listaMadrugada []int

	for _, v := range ListadoRecuperadoTickets {
		hora := strings.Split(v.HoraVuelo, ":")
		horaInt, err := strconv.Atoi(hora[0])
		if err != nil {
			panic("Error de conversión de dato")
		}
		switch {
		case horaInt >= 0 && horaInt <= 6:
			listaMadrugada = append(listaMadrugada, horaInt)

		case horaInt >= 7 && horaInt <= 12:
			listaMañana = append(listaMañana, horaInt)
		case horaInt >= 13 && horaInt <= 19:
			listaTarde = append(listaTarde, horaInt)

		case horaInt >= 20 && horaInt <= 23:
			listaNoche = append(listaNoche, horaInt)
		}

	}
	var total int
	switch time {
	case "Madrugada":
		total = len(listaMadrugada)
		fmt.Printf("La cantidad total de tickets para la madrugada es %d ", total)
		return total, nil
	case "Mañana":
		total = len(listaMañana)
		fmt.Printf("La cantidad total de tickets para la mañana es %d", total)
		return total, nil
	case "Tarde":
		total = len(listaTarde)
		fmt.Printf("La cantidad total de tickets para la tarde es %d", total)
		return total, nil
	case "Noche":
		total = len(listaNoche)
		fmt.Printf("La cantidad total de tickets para la noche es %d", total)
		return total, nil
	default:
		fmt.Println(e)
		return 0, e

	}

}

// Función para obtener porcentaje según destino
func ObtenerPromedioDestinos(destino string) (float64, error) {
	ListadoRecuperadoTickets, err := ObtenerDatos(archivo)
	if err != nil {
		fmt.Println(err)
	}
	e := errors.New("Error en el listado")
	totalListado := float64(len(ListadoRecuperadoTickets))
	totalDestinos, er := ObtenerTotalTicketsDestino(destino)
	parseTotalDestinos := float64(totalDestinos)
	if er != nil {
		fmt.Println(er)
	}
	if totalListado == 0 {
		fmt.Println(e)
		return 0, e
	}
	porcentaje := (parseTotalDestinos * 100) / totalListado
	fmt.Printf("El porcentaje total de tickets para el destino %s es %.2f", destino, porcentaje)

	return porcentaje, nil
}
