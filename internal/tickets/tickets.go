package tickets

import (
	"errors"
	"fmt"
	"log"
	"os"
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

var ListadoRecuperado *[]Ticket

// TODO: Consultar si es necesario cerrar archivo
// TODO: Falta el defer
func ObtenerDatos(ruta string) (ListadoRecuperado *[]Ticket, e error) {
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
		panic("Error de lectura") //return nil, err
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

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	var Listado *[]Ticket
	Listado, _ = ObtenerDatos("tickets.csv")
	acum := 0

	for _, v := range *Listado {
		if v.PaisDestino == destination {
			acum++
		}
	}
	if acum == 0 {
		return 0, errors.New("No se encontraron coincidencias con el destino")
	}
	return acum, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	return 0, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 0, nil
}
