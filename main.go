package main

import (
	"fmt"
	"log"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
)

func main() {
	var pais string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&pais)
	totalTickets, e := tickets.GetTotalTickets(pais)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(totalTickets)
	var franjaHoraria string
	fmt.Println("Ingrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	fmt.Scan(&franjaHoraria)
	totalVuelosTime, e := tickets.GetTime(franjaHoraria)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(totalVuelosTime)
}
