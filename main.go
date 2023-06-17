package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
)

func main() {
	var archivo string = "tickets.csv"
	Listado, e := tickets.ObtenerDatos(archivo)
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

	//GO ROUTINE 1 (Ejemplo de invocación de una función dentro de una función anónima)
	go func(p string, a []tickets.Ticket) {
		total, err := tickets.ObtenerTotalTicketsDestino(destino, Listado)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\nLa cantidad total de tickets para %s es %d", destino, total)
	}(destino, Listado)

	//GO ROUTINE 2 (Ejemplo de routine invocando la función con la paralabra reservada "go")
	go func(f string, a []tickets.Ticket) {
		_, err := tickets.ObtenerTicketsFranjaHoraria(franjaHoraria, Listado)
		if err != nil {
			fmt.Println(err)
			return
		}
	}(franjaHoraria, Listado)

	//GO ROUTINE 3
	go func(d string, a []tickets.Ticket) {
		_, err := tickets.ObtenerPromedioDestinos(destinoPorcentaje, Listado)
		if err != nil {
			fmt.Println(err)
			return
		}

	}(destinoPorcentaje, Listado)

	//FIXME, NO IMPRIME NADA SIN EL TIME SLEEP
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
