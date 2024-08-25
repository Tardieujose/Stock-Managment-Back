package auth

import (
    "encoding/json"
	"log"
    "net/http"
    "github.com/Tardieujose/Stock-Managment-Back/internal/user"
	"github.com/Tardieujose/Stock-Managment-Back/pkg/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    // Decodifica el JSON recibido en la solicitud
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        log.Println("Error decoding request payload:", err)
        return
    }

    // Busca el usuario en la base de datos
    usr, err := user.FindUserByUsername(creds.Username)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        log.Println("Error finding user:", err)
        return
    }
    if usr == nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        log.Println("User not found")
        return
    }

    // Compara la contraseña
    if usr.Password != creds.Password {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        log.Println("Password mismatch")
        return
    }

    // Genera el token JWT
    token, err := utils.GenerateJWT(usr.Username, usr.Role)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        log.Println("Error generating JWT:", err)
        return
    }

    // Devuelve el token JWT en la respuesta
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
