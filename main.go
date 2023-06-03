package main

import (
	"fmt"
	"time"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
)

func main() {

	fmt.Println("Bienvenido a GolangAirlines")

	var pais string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&pais)

	//GO ROUTINE 1 (Ejemplo de invocacion de una función dentro de una función anónima)
	go func(p string) {
		total, err := tickets.GetTotalTickets(pais)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("La cantidad total de tickets para %s es %d", pais, total)
	}(pais)

	time.Sleep(1 * time.Second)

	var franjaHoraria string
	fmt.Println("\nIngrese una franja horaria (Disponibles: Madrugada, Mañana, Tarde, Noche):")
	fmt.Scan(&franjaHoraria)

	//GO ROUTINE 2 (Ejemplo de routine invocando la función con la paralabra reservada "go")
	go tickets.GetTime(franjaHoraria)
	time.Sleep(1 * time.Second)

	var destinoPorcentaje string
	fmt.Println("\nDigite el destino para saber el porcentaje de pasajeros que viajo en el dia")
	fmt.Scan(&destinoPorcentaje)

	//GO ROUTINE 3
	go func(d string) {
		tickets.AverageDestination(destinoPorcentaje)
	}(destinoPorcentaje)
	time.Sleep(1 * time.Second)

	// SI QUIERE VER EL LISTADO COMPLETO OBTENIDO DESCOMENTE ESTA FUNCIÓN
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
