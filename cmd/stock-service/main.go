package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    // "github.com/rs/cors"
    "github.com/Tardieujose/Stock-Managment-Back/internal/stock"  // Ruta al paquete product para los handlers
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/stock", stock.GetStockHandler).Methods("GET")
    router.HandleFunc("/stock", stock.UpdateStockHandler).Methods("PUT")

    log.Println("Stock service is running on port 8002")
    log.Fatal(http.ListenAndServe(":8002", router))
}