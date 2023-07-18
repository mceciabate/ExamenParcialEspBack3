package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
)

func main() {
	var archivo string = "tickets.csv"
	listado, e := tickets.ObtenerDatos(archivo)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("***BIENVENIDO A GOLANGAIRLINES***")

	var destino string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&destino)
	var franjaHoraria string
	fmt.Println("\nIngrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	fmt.Scan(&franjaHoraria)
	var destinoPorcentaje string
	fmt.Println("\nDigite el destino para saber el porcentaje de pasajeros que viajó en el día")
	fmt.Scan(&destinoPorcentaje)

	//GO ROUTINE 1
	go func(p string, a []tickets.Ticket) {
		total, err := tickets.ObtenerTotalTicketsDestino(destino, &listado)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\nLa cantidad total de tickets para %s es %d", destino, total)
	}(destino, listado)

	//GO ROUTINE 2
	go func(f string, a []tickets.Ticket) {
		total, err := tickets.ObtenerTicketsFranjaHoraria(franjaHoraria, &listado)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\nLa cantidad total de tickets para la %s es %d\n", f, total)

	}(franjaHoraria, listado)

	//GO ROUTINE 3
	go func(d string, a []tickets.Ticket) {
		porcentaje, err := tickets.ObtenerPromedioDestinos(destinoPorcentaje, &listado)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\nEl porcentaje total de tickets para el destino %s es %.2f", d, porcentaje)

	}(destinoPorcentaje, listado)

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
