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

	// var franjaHoraria string
	// fmt.Println("Ingrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	// fmt.Scan(&franjaHoraria)
	// totalVuelosTime, e := tickets.GetTime(franjaHoraria)
	// if e != nil {
	// 	fmt.Println(e)
	// }
	// fmt.Println(totalVuelosTime)
	// var destinoPorcentaje string
	// fmt.Println("Digite el destino para saber el porcentaje de pasajeros que viajo en el dia")
	// fmt.Scan(&destinoPorcentaje)
	// porcentajeDestino, err := tickets.AverageDestination(destinoPorcentaje)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(porcentajeDestino)

}
