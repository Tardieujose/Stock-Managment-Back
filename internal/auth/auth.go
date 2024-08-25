package auth

import (
    "github.com/gorilla/mux"
)

// RegisterRoutes configura las rutas relacionadas con la autenticación
func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/api/login", LoginHandler).Methods("POST")
}
