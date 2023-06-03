package main

import (
	"fmt"
	"time"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
)

func main() {

	fmt.Println("***BIENVENIDO A GOLANGAIRLINES***")

	var destino string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&destino)

	//GO ROUTINE 1 (Ejemplo de invocación de una función dentro de una función anónima)
	go func(p string) {
		total, err := tickets.ObtenerTotalTicketsDestino(destino)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("La cantidad total de tickets para %s es %d", destino, total)
	}(destino)

	time.Sleep(1 * time.Second)

	var franjaHoraria string
	fmt.Println("\nIngrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	fmt.Scan(&franjaHoraria)

	//GO ROUTINE 2 (Ejemplo de routine invocando la función con la paralabra reservada "go")
	go tickets.ObtenerTicketsFranjaHoraria(franjaHoraria)
	time.Sleep(1 * time.Second)

	var destinoPorcentaje string
	fmt.Println("\nDigite el destino para saber el porcentaje de pasajeros que viajó en el día")
	fmt.Scan(&destinoPorcentaje)

	//GO ROUTINE 3
	go func(d string) {
		tickets.ObtenerPromedioDestinos(destinoPorcentaje)
	}(destinoPorcentaje)
	time.Sleep(1 * time.Second)

	fmt.Println("\n***GRACIAS POR USAR EL SERVICIO DE GOLANGAIRLINES***")

	// SI QUIERE VER EL LISTADO COMPLETO OBTENIDO DESCOMENTE ESTA SECCIÓN
	// go func() {
	// 	Listado, e := tickets.ObtenerDatos("tickets.csv")
	// 	if e != nil {
	// 		fmt.Println(e)
	// 	}
	// 	fmt.Println("*******LISTADO COMPLETO DE TICKETS**********")
	// 	fmt.Println(Listado)
	// }()
	// time.Sleep(1 * time.Second)

}
