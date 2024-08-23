package main

import (
    "log"
    "net/http"

    "github.com/Tardieujose/Stock-Managment-Back/pkg/database"
    "github.com/gorilla/mux"
    // "github.com/rs/cors"
    "github.com/Tardieujose/Stock-Managment-Back/internal/product"  // Ruta al paquete product para los handlers
)

func main() {

    database.ConnectDB()
    router := mux.NewRouter()

    router.HandleFunc("/products", product.GetProductsHandler).Methods("GET")
    router.HandleFunc("/products", product.CreateProductHandler).Methods("POST")

    log.Println("Product service is running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
