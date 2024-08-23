package main

import (
	"log"
	"net/http"

	"github.com/Tardieujose/Stock-Managment-Back/pkg/database"
	"github.com/gorilla/mux"
	// "github.com/rs/cors"
	"github.com/Tardieujose/Stock-Managment-Back/internal/order" // Ruta al paquete product para los handlers
)

func main() {
	
    database.ConnectDB()
    router := mux.NewRouter()

	router.HandleFunc("/orders", order.GetOrdersHandler).Methods("GET")
	router.HandleFunc("/orders", order.CreateOrderHandler).Methods("POST")

	log.Println("Order service is running on port 8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}
