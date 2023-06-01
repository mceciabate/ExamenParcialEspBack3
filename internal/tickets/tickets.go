package tickets

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// id, nombre, email, país de destino, hora del vuelo y precio.
type Ticket struct {
	Id             int64
	NombreCompleto string
	Email          string
	PaisDestino    string
	HoraVuelo      string
	Precio         int64
}

var ListadoRecuperado []Ticket

// TODO: Consultar si es necesario cerrar archivo
// TODO: Falta el defer
func ObtenerDatos(ruta string) (a []Ticket, e error) {
	var array []Ticket
	var newTicket Ticket
	var line4 string
	var newArray [][]string
	var otroArray []string
	f, err := os.Open(ruta)
	if err != nil {
		panic(err)
	}
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
			id, err := strconv.ParseInt(otroArray[0], 16, 64)
			if err != nil {
				log.Fatal("Error de conversión")
			}
			precio, err1 := strconv.ParseInt(otroArray[5], 16, 64)
			if err1 != nil {
				log.Fatal("Error de conversión")
			}
			newTicket = Ticket{
				Id:             id,
				NombreCompleto: otroArray[1],
				Email:          otroArray[2],
				PaisDestino:    otroArray[3],
				HoraVuelo:      otroArray[4],
				Precio:         precio,
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

	return nil, nil
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	return 0, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	return 0, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 0, nil
}
