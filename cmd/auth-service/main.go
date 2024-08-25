package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/database"
    "github.com/Tardieujose/Stock-Managment-Back/internal/auth"
)

func main() {
    database.ConnectDB()

    r := mux.NewRouter()
    auth.RegisterRoutes(r) // Usa RegisterRoutes para definir tus rutas de autenticación
    
    // Configura CORS
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Permite el origen del frontend
        handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"}), // Permite métodos HTTP
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Permite encabezados específicos
    )(r)

    log.Println("Auth service is running on port 8003")
    log.Fatal(http.ListenAndServe(":8003", corsHandler))
}
