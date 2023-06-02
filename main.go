package main

import (
	"fmt"
	"log"

	"github.com/mceciabate/ExamenParcialEspBack3/internal/tickets"
	// "github.com/mceciabate/ExamenParcialEspBack3/tickets"
)

func main() {
	// total, err := tickets.GetTotalTickets("Brazil")
	// lista, _ := tickets.ObtenerDatos("tickets.csv")
	// fmt.Println(lista,"----------------------")
	fmt.Println(tickets.ListadoRecuperadoTickets)
	var pais string
	fmt.Println("Por favor ingrese un destino: ")
	fmt.Scan(&pais)
	totalTickets, e := tickets.GetTotalTickets(pais)
	if e != nil {
		log.Fatal(e)
	}

	// totalTickets, _:=	tickets.GetTotalTickets("Colombia")
	fmt.Println(totalTickets)
}
