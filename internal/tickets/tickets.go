package tickets

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// id, nombre, email, país de destino, hora del vuelo y precio.
type Ticket struct {
	Id             string
	NombreCompleto string
	Email          string
	PaisDestino    string
	HoraVuelo      string
	Precio         string
}

// TODO TRATAR ERROR
var ListadoRecuperadoTickets, e = ObtenerDatos("tickets.csv")

// TODO: Consultar si es necesario cerrar archivo
// TODO: Falta el defer

// Funcion para obtener datos
func ObtenerDatos(ruta string) (a *[]Ticket, e error) {
	var array []Ticket
	var newTicket Ticket
	var line4 string
	var newArray [][]string
	var otroArray []string
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se produjo un error abriendo el archivo")
		}
	}()
	f, err := os.Open(ruta)
	if err != nil {
		panic("Error de lectura")
	}
	defer f.Close()
	rawData, er := os.ReadFile(ruta)
	if er != nil {
		panic("Error de lectura de archivo")
	}
	data := strings.Split(string(rawData), "; ")
	for _, v := range data {
		line := strings.Split(v, "\n")
		for _, v := range line {
			line2 := strings.Split(v, " ")
			newArray = append(newArray, line2)
		}
		for i := 0; i < len(newArray); i++ {
			line3 := newArray[i]
			line4 = strings.Join(line3, " ")
			otroArray = strings.Split(line4, ",")
			// id, err := strconv.ParseInt(otroArray[0], 16, 64)
			// if err != nil {
			// 	log.Fatal("Error de conversión")
			// }
			// precio, err1 := strconv.ParseInt(otroArray[5], 16, 64)
			// if err1 != nil {
			// 	log.Fatal("Error de conversión")
			// }
			newTicket = Ticket{
				Id:             otroArray[0],
				NombreCompleto: otroArray[1],
				Email:          otroArray[2],
				PaisDestino:    otroArray[3],
				HoraVuelo:      otroArray[4],
				Precio:         otroArray[5],
			}
			array = append(array, newTicket)
		}
		if len(array) == 0 {
			return nil, errors.New("No se ha generado el listado de forma correcta")
		}
	}
	err5 := f.Close()
	if err5 != nil {
		log.Fatal("No se puede cerrar el archivo")
	}

	return &array, nil
}

// Funcion para obtener el listado de Tickets según destino
func GetTotalTickets(destination string) (int, error) {
	acum := 0
	for _, v := range *ListadoRecuperadoTickets {
		if v.PaisDestino == destination {
			acum++
		}
	}
	if acum == 0 {
		return 0, errors.New("No se encontraron coincidencias con el destino")
	}
	return acum, nil
}

// Función para obtener Tickets segun franja horaria
func GetTime(time string) (int, error) {

	var listaMañana []int
	var listaTarde []int
	var listaNoche []int
	var listaMadrugada []int

	for _, v := range *ListadoRecuperadoTickets {
		hora := strings.Split(v.HoraVuelo, ":")
		horaInt, err := strconv.Atoi(hora[0])
		if err != nil {
			log.Fatal(err)
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
	switch time {
	case "Madrugada":
		return len(listaMadrugada), nil
	case "Mañana":
		return len(listaMañana), nil
	case "Tarde":
		return len(listaTarde), nil
	case "Noche":
		return len(listaNoche), nil
	default:
		return 0, errors.New("Ingrese una franja horaria válida")

	}

}

// Función para obtener porcentaje segun destino
func AverageDestination(destination string) (float64, error) {
	totalListado := float64(len(*ListadoRecuperadoTickets))
	totalDestinos, err := GetTotalTickets(destination)
	parseTotalDestinos := float64(totalDestinos)
	if err != nil {
		fmt.Println(err)
	}
	if totalListado == 0 {
		return 0, errors.New("Error en el listado ")
	}
	porcentaje := (parseTotalDestinos * 100) / totalListado
	return porcentaje, nil
}
