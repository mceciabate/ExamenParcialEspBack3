package main

import (
	"fmt"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
)

func main() {
	var pais string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&pais)
	totalTickets, e := tickets.GetTotalTickets(pais)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(totalTickets)
	var franjaHoraria string
	fmt.Println("Ingrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	fmt.Scan(&franjaHoraria)
	totalVuelosTime, e := tickets.GetTime(franjaHoraria)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(totalVuelosTime)
	var destinoPorcentaje string
	fmt.Println("Digite el destino para saber el porcentaje de pasajeros que viajo en el dia")
	fmt.Scan(&destinoPorcentaje)
	porcentajeDestino, err := tickets.AverageDestination(destinoPorcentaje)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(porcentajeDestino)

}
