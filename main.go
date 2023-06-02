package main

import (
	"fmt"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
	// "github.com/mceciabate/ExamenParcialEspBack3/tickets"
)

func main() {
	// total, err := tickets.GetTotalTickets("Brazil")
	// lista, _ := tickets.ObtenerDatos("tickets.csv")
	// fmt.Println(lista,"----------------------")
	var pais string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&pais)
	totalTickets, _ := tickets.GetTotalTickets(pais)

	// totalTickets, _:=	tickets.GetTotalTickets("Colombia")
	fmt.Println(totalTickets)
}
