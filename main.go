package main

import (
	"fmt"
	"time"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
)

func main() {

	var pais string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&pais)

	go func(p string) {

		tickets.GetTotalTickets(pais)

	}(pais)
	time.Sleep(1 * time.Second)

	var franjaHoraria string
	fmt.Println("\nIngrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	fmt.Scan(&franjaHoraria)
	go func (f string)  {
		tickets.GetTime(franjaHoraria)
	}(franjaHoraria)
	time.Sleep(1 * time.Second)
	
	var destinoPorcentaje string
	fmt.Println("\nDigite el destino para saber el porcentaje de pasajeros que viajo en el dia")
	fmt.Scan(&destinoPorcentaje)

	go func(d string) {
		tickets.AverageDestination(destinoPorcentaje)
	}(destinoPorcentaje)
	time.Sleep(1 * time.Second)
	
}
