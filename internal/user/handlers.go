package user

import (
    "encoding/json"
    "net/http"
    // "github.com/Tardieujose/Stock-Managment-Back/pkg/utils" // Utiliza las funciones de utilidad para la validación de JWT
)

// RegisterHandler maneja el registro de nuevos usuarios
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var newUser User
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    // Aquí podrías agregar lógica para encriptar la contraseña antes de guardar el usuario
    _, err := CreateUser(newUser)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User registered successfully"))
}

// GetUserHandler maneja la obtención de un usuario por su nombre de usuario
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    if username == "" {
        http.Error(w, "Username is required", http.StatusBadRequest)
        return
    }

    user, err := FindUserByUsername(username)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    response, err := json.Marshal(user)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}
