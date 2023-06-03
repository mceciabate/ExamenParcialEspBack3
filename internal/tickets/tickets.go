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

// TODO: Consultar si es necesario cerrar archivo
// TODO: Falta el defer

// Funcion para obtener datos
func ObtenerDatos(ruta string) ([]Ticket, error) {
	var array []Ticket
	var newTicket Ticket
	var line4 string
	var arrayStrings [][]string
	var arrayTicket []string
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
	err5 := f.Close()
	if err5 != nil {
		log.Fatal("No se puede cerrar el archivo")
	}
	return array, nil
}

// Funcion para obtener el listado de Tickets según destino
func GetTotalTickets(destino string) (int, error) {
	ListadoRecuperadoTickets, err := ObtenerDatos("tickets.csv")
	if err != nil {
		panic("No se puede obtener el listado")
	}
	e := errors.New("No se encontraron coincidencias con el destino")
	acum := 0
	for _, v := range ListadoRecuperadoTickets {
		if v.PaisDestino == destino {
			acum++
		}
	}
	if acum == 0 {
		// fmt.Println(e)
		return 0, e
	}
	// fmt.Printf("La cantidad total de tickets para %s es %d", destino, acum)
	return acum, nil
}

// Función para obtener Tickets segun franja horaria
func GetTime(time string) (int, error) {
	ListadoRecuperadoTickets, err := ObtenerDatos("tickets.csv")
	if err != nil {
		panic("No se puede obtener el listado")
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
			fmt.Println(err)
			// log.Fatal(err)
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

// Función para obtener porcentaje segun destino
func AverageDestination(destino string) (float64, error) {
	ListadoRecuperadoTickets, err := ObtenerDatos("tickets.csv")
	if err != nil {
		panic("No se puede obtener el listado")
	}
	e := errors.New("Error en el listado ")
	totalListado := float64(len(ListadoRecuperadoTickets))
	totalDestinos, er := GetTotalTickets(destino)
	parseTotalDestinos := float64(totalDestinos)
	if er != nil {
		log.Fatal(er)
	}
	if totalListado == 0 {
		fmt.Println(e)
		return 0, e
	}
	porcentaje := (parseTotalDestinos * 100) / totalListado
	fmt.Printf("El porcentaje total de tickets para destino %s es %.2f", destino, porcentaje)

	return porcentaje, nil
}
