package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/database"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/utils" // Cambia la ruta a pkg/utils para JWT
    "github.com/Tardieujose/Stock-Managment-Back/internal/product"  // Ruta al paquete product para los handlers
)

func main() {
    database.ConnectDB()  // Asegúrate de tener una función correcta para conectar con MongoDB
    router := mux.NewRouter()

    // Ruta pública (sin autenticación)
    router.HandleFunc("/products", product.GetProductsHandler).Methods("GET")

    // Ruta protegida (requiere autenticación y rol admin)
    router.HandleFunc("/products", authenticateAdmin(product.CreateProductHandler)).Methods("POST")

    log.Println("Product service is running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}

// Middleware para autenticar las solicitudes
func authenticateAdmin(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Missing authorization token", http.StatusUnauthorized)
            return
        }

        // Elimina el prefijo "Bearer " si está presente
        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        }

        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Verifica si el rol es admin
        if claims.Role != "admin" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Llama al siguiente handler
        next.ServeHTTP(w, r)
    }
}
